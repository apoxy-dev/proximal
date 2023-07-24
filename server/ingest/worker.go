package ingest

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/activity"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"golang.org/x/mod/sumdb/dirhash"

	"github.com/apoxy-dev/proximal/server/build"
	gobuild "github.com/apoxy-dev/proximal/server/build/go"
	rustbuild "github.com/apoxy-dev/proximal/server/build/rust"
	serverdb "github.com/apoxy-dev/proximal/server/db"
	sqlc "github.com/apoxy-dev/proximal/server/db/sql"
	"github.com/apoxy-dev/proximal/server/watcher"

	middlewarev1 "github.com/apoxy-dev/proximal/api/middleware/v1"
)

const (
	MiddlewareIngestQueue = "MIDDLEWARE_INGEST_TASK_QUEUE"
)

type MiddlewareIngestParams struct {
	Slug string

	// Temporal can't unmarshal oneofs in protos so use bytes and unmarshal ourselves.
	Params *middlewarev1.MiddlewareIngestParams
}

type MiddlewareIngestResult struct {
	// SHA256 of the build.
	// For local builds this is the hash of the build directory + params.
	// For GIT builds this is the hash of the commit + params.
	SHA string

	// Whether the build was previously cached (a build with SHA above exists in the store).
	Cached bool

	// Err is set if build fails midway.
	Err string
}

// StartBuildWorkflow spawns another async workflow to perform actual build and releases,
// and returns immediately to release the API call.
func StartBuildWorkflow(ctx workflow.Context, in *MiddlewareIngestParams) error {
	cwo := workflow.ChildWorkflowOptions{
		WorkflowExecutionTimeout: 30 * time.Minute,
		// Continue running child workflow on parent completion.
		ParentClosePolicy: enumspb.PARENT_CLOSE_POLICY_ABANDON,
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	future := workflow.ExecuteChildWorkflow(ctx, DoBuildWorkflow, in)
	// NB: This needs to be called before parent exits or child workflow will be lost.
	// See: https://github.com/temporalio/temporal/issues/685
	if err := future.GetChildWorkflowExecution().Get(ctx, nil); err != nil {
		workflow.GetLogger(ctx).Error("DoBuildWorkflow failed", "Error", err)
		return err
	}
	return nil
}

// DoBuildWorkflow is a workflow executes build Activities. Meant to be run async from a parent
// workflow that quits to release the API call.
func DoBuildWorkflow(ctx workflow.Context, in *MiddlewareIngestParams) error {
	opts := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute, // 10 min to download WASM file should be enough?
		// HeartbeatTimeout:    2 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        time.Second,
			BackoffCoefficient:     2.0,
			MaximumInterval:        60 * time.Second,
			MaximumAttempts:        3,
			NonRetryableErrorTypes: []string{"FatalError"},
		},
	}
	ctx = workflow.WithActivityOptions(ctx, opts)
	logger := tlog.With(workflow.GetLogger(ctx), "Slug", in.Slug)

	so := &workflow.SessionOptions{
		CreationTimeout:  time.Minute,
		ExecutionTimeout: 5 * time.Minute,
	}
	sessCtx, err := workflow.CreateSession(ctx, so)
	if err != nil {
		return err
	}
	defer workflow.CompleteSession(sessCtx)

	var w *IngestWorker // Using nil worker for activity triggers.

	var buildResult *MiddlewareIngestResult
	if in.Params.Type == middlewarev1.MiddlewareIngestParams_GITHUB {
		err = workflow.ExecuteActivity(sessCtx, w.PrepareGithubBuildActivity, in).Get(ctx, &buildResult)
	} else { // Direct.
		err = workflow.ExecuteActivity(sessCtx, w.PrepareLocalBuildActivity, in).Get(ctx, &buildResult)
	}
	if err != nil {
		logger.Error("Prepare activity failed", "Error", err)
		buildResult = &MiddlewareIngestResult{
			Err: err.Error(),
		}
		finErr := workflow.ExecuteActivity(sessCtx, w.FinalizeActivity, in, buildResult).Get(ctx, nil)
		if finErr != nil {
			logger.Error("FinalizeActivity failed", "Error", finErr)
		}
		return err
	}

	if !buildResult.Cached {
		logger.Info("Build already cached", "SHA", buildResult.SHA)

		if err = workflow.ExecuteActivity(sessCtx, w.BuildActivity, buildResult.SHA, in).Get(ctx, nil); err != nil {
			logger.Error("BuildActivity failed", "Error", err)

			buildResult.Err = err.Error()
			finErr := workflow.ExecuteActivity(sessCtx, w.FinalizeActivity, in, buildResult).Get(ctx, nil)
			if finErr != nil {
				logger.Error("FinalizeActivity failed", "Error", finErr)
			}
			return err
		}

		if err = workflow.ExecuteActivity(sessCtx, w.UploadWasmOutputActivity, in.Slug, buildResult.SHA).
			Get(ctx, nil); err != nil {

			logger.Error("UploadWasmOutputActivity failed", "Error", err)

			buildResult.Err = err.Error()
			finErr := workflow.ExecuteActivity(sessCtx, w.FinalizeActivity, in, buildResult).Get(ctx, nil)
			if finErr != nil {
				logger.Error("FinalizeActivity failed", "Error", finErr)
			}
			return err
		}
	}

	if finErr := workflow.ExecuteActivity(sessCtx, w.FinalizeActivity, in, buildResult).Get(ctx, nil); finErr != nil {
		logger.Error("FinalizeActivity failed", "Error", finErr)
		return finErr
	}

	logger.Info("Build complete", "SHA", buildResult.SHA)

	return nil
}

// IngestWorker runs build activities.
type IngestWorker struct {
	workDir       string
	remoteStorage bool
	db            *serverdb.DB
	watcher       *watcher.Watcher
}

func NewIngestWorker(workDir string, db *serverdb.DB, w *watcher.Watcher) *IngestWorker {
	return &IngestWorker{
		workDir: workDir,
		db:      db,
		watcher: w,
	}
}

func copyRunDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		}
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}
		// Link and ignore if file already exists.
		if err := forceCopyFile(path, dstPath); err != nil {
			return err
		}
		return nil
	})
}

func forceCopyFile(source, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(destination)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			os.Remove(destination)
			if dst, err = os.Create(destination); err != nil {
				return err
			}
		}
		return err
	}
	_, err = io.Copy(dst, src)
	dst.Close()
	if err != nil {
		return err
	}
	fi, err := os.Stat(source)
	if err != nil {
		os.Remove(destination)
		return err
	}
	err = os.Chmod(destination, fi.Mode())
	if err != nil {
		os.Remove(destination)
		return err
	}
	return nil
}

// hash is similar to https://cs.opensource.google/go/x/mod/+/refs/tags/v0.12.0:sumdb/dirhash/hash.go;drc=b71060237c896c7c9fc602ab66e33ea6079659fa;l=44
// but doesn't base64 encode the output.
func hash(files []string, open func(string) (io.ReadCloser, error)) (string, error) {
	h := sha256.New()
	files = append([]string(nil), files...)
	sort.Strings(files)
	for _, file := range files {
		if strings.Contains(file, "\n") {
			return "", errors.New("dirhash: filenames with newlines are not supported")
		}
		r, err := open(file)
		if err != nil {
			return "", err
		}
		hf := sha256.New()
		_, err = io.Copy(hf, r)
		r.Close()
		if err != nil {
			return "", err
		}
		fmt.Fprintf(h, "%x  %s\n", hf.Sum(nil), file)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func (w *IngestWorker) getTargetSHA(
	ctx context.Context,
	slug string,
	params *middlewarev1.MiddlewareIngestParams,
) (sha string, err error) {
	if params == nil {
		return "", errors.New("params cannot be nil")
	}

	switch params.Type {
	case middlewarev1.MiddlewareIngestParams_DIRECT:
		sha, err = dirhash.HashDir(params.WatchDir, "", hash)
		if err != nil {
			return "", fmt.Errorf("failed to hash directory: %v", err)
		}
	default:
		return "", fmt.Errorf("unknown ingest type: %v", params.Type)
	}
	return sha, nil
}

func (w *IngestWorker) checkIfBuildExists(ctx context.Context, slug string, sha string) (*sqlc.Build, bool, error) {
	logger := tlog.With(activity.GetLogger(ctx), "Slug", slug, "Sha", sha)
	b, err := w.db.Queries().GetBuildByMiddlewareSlugAndSha(ctx, sqlc.GetBuildByMiddlewareSlugAndShaParams{
		MiddlewareSlug: slug,
		Sha:            sha,
	})
	if err == nil {
		logger.Info("Build already exists")

		wasmOut := w.wasmOut(ctx, slug, sha)
		if _, err := os.Stat(wasmOut); err == nil {
			logger.Info("Build output already exists", "WasmOut", wasmOut)
			return &b, true, nil
		}

		logger.Warn("Build exists in DB but not in store, re-building")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return nil, false, fmt.Errorf("failed to get build: %w", err)
	}

	return nil, false, nil
}

func (w *IngestWorker) buildDir(slug string, sha string) string {
	return filepath.Join(w.workDir, slug, sha)
}

func (w *IngestWorker) runDir(ctx context.Context, slug string, sha string) string {
	return filepath.Join(w.buildDir(slug, sha), activity.GetInfo(ctx).WorkflowExecution.RunID)
}

func (w *IngestWorker) srcDir(ctx context.Context, slug string, sha string) string {
	return filepath.Join(w.runDir(ctx, slug, sha), "src")
}

func (w *IngestWorker) wasmOut(ctx context.Context, slug string, sha string) string {
	return filepath.Join(w.buildDir(slug, sha), "wasm.out")
}

func (w *IngestWorker) PrepareGithubBuildActivity(
	ctx context.Context,
	in *MiddlewareIngestParams,
) (*MiddlewareIngestResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("PrepareGithubBuildActivity", "Slug", in.Slug)

	commit := in.Params.Commit
	if commit == "" {
		r := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
			Name: "origin",
			URLs: []string{in.Params.GithubRepo},
		})
		ref, err := r.ListContext(ctx, &git.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to list remote: %w", err)
		}
		for _, r := range ref {
			if (r.Name().IsBranch() || r.Name().IsTag()) && r.Name().Short() == in.Params.Branch {
				commit = r.Hash().String()
				break
			}
		}
		if commit == "" {
			return nil, fmt.Errorf("failed to find commit for ref %q", in.Params.Branch)
		}
	}

	b, exists, err := w.checkIfBuildExists(ctx, in.Slug, commit)
	if err != nil {
		return nil, fmt.Errorf("failed to check if build exists: %w", err)
	}
	if exists {
		return &MiddlewareIngestResult{
			SHA:    b.Sha,
			Cached: true,
		}, nil
	}

	// TODO(dilyevsky): If src dir exists, we should just pull instead of clone.
	if err := os.RemoveAll(w.srcDir(ctx, in.Slug, commit)); err != nil {
		return nil, fmt.Errorf("failed to remove old src dir: %w", err)
	}
	if err := os.MkdirAll(w.buildDir(in.Slug, commit), 0755); err != nil {
		return nil, fmt.Errorf("failed to create run dir: %w", err)
	}

	logger.Info("Cloning GitHub repo", "Repo", in.Params.GithubRepo, "Commit", commit)

	repo, err := git.PlainCloneContext(
		ctx, w.srcDir(ctx, in.Slug, commit),
		false, /* bare */
		&git.CloneOptions{
			URL:      in.Params.GithubRepo,
			Progress: os.Stdout,
		})
	if err != nil {
		return nil, fmt.Errorf("failed to clone repo: %w", err)
	}
	wTree, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get worktree: %w", err)
	}
	cOpts := &git.CheckoutOptions{}
	if commit != "" {
		cOpts.Hash = plumbing.NewHash(commit)
	} else {
		cOpts.Branch = plumbing.NewBranchReferenceName(in.Params.Branch)
	}
	if err := wTree.Checkout(cOpts); err != nil {
		return nil, fmt.Errorf("failed to checkout commit: %w", err)
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get head: %w", err)
	}
	sha := ref.Hash().String()

	_, err = w.db.Queries().CreateBuild(ctx, sqlc.CreateBuildParams{
		Sha:            sha,
		MiddlewareSlug: in.Slug,
		Status:         middlewarev1.Build_PREPARING.String(),
		StatusDetail:   "Preparing GitHub build",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create build: %w", err)
	}

	return &MiddlewareIngestResult{
		SHA:    sha,
		Cached: false,
	}, nil
}

func (w *IngestWorker) PrepareLocalBuildActivity(
	ctx context.Context,
	in *MiddlewareIngestParams,
) (*MiddlewareIngestResult, error) {
	logger := tlog.With(activity.GetLogger(ctx), "Slug", in.Slug)
	logger.Info("PrepareLocalBuildActivity")

	sha, err := w.getTargetSHA(ctx, in.Slug, in.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to get target sha: %v", err)
	}
	logger = tlog.With(logger, "SHA", sha)

	b, exists, err := w.checkIfBuildExists(ctx, in.Slug, sha)
	if err != nil {
		return nil, fmt.Errorf("failed to check if build exists: %w", err)
	}
	if exists {
		return &MiddlewareIngestResult{
			SHA:    b.Sha,
			Cached: true,
		}, nil
	}

	srcDir := w.srcDir(ctx, in.Slug, sha)
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create run dir: %w", err)
	}
	if err := copyRunDir(in.Params.WatchDir, srcDir); err != nil {
		return nil, fmt.Errorf("failed to copy watch dir: %w", err)
	}

	_, err = w.db.Queries().CreateBuild(ctx, sqlc.CreateBuildParams{
		Sha:            sha,
		MiddlewareSlug: in.Slug,
		Status:         middlewarev1.Build_PREPARING.String(),
		StatusDetail:   "Preparing initial local build",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create build: %v", err)
	}

	logger.Info("Created build")

	return &MiddlewareIngestResult{
		SHA:    sha,
		Cached: false,
	}, nil
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func tmpSuffix() string {
	rs := make([]rune, 7)
	for i := range rs {
		rs[i] = chars[rand.Intn(len(chars))]
	}
	return string(rs)
}

func copyLogs(dir string, file string, r io.Reader, attempt int) error {
	f, err := os.Create(filepath.Join(dir, fmt.Sprintf("%s.%d", file, attempt)))
	if err != nil {
		return err
	}
	defer f.Close()

	// stdout -> stdout.1 (via tmp file for atomicity)
	tmp := fmt.Sprintf("/tmp/%s.%s", file, tmpSuffix())
	if err := os.Symlink(f.Name(), tmp); err != nil {
		return fmt.Errorf("failed to symlink: ", err)
	}
	if err := os.Rename(tmp, filepath.Join(dir, file)); err != nil {
		return fmt.Errorf("failed to rename: ", err)
	}

	if _, err := io.Copy(f, r); err != nil {
		return fmt.Errorf("failed to copy stream: ", err)
	}
	return nil
}

func (w *IngestWorker) BuildActivity(ctx context.Context, sha string, in *MiddlewareIngestParams) error {
	logger := activity.GetLogger(ctx)
	logger.Info("BuildActivity", "Slug", in.Slug)

	var builder build.Builder
	switch in.Params.Language {
	case middlewarev1.MiddlewareIngestParams_GO:
		builder = gobuild.NewBuilder()
	case middlewarev1.MiddlewareIngestParams_RUST:
		builder = rustbuild.NewBuilder()
	default:
		return fmt.Errorf("unsupported language: %v", in.Params.Language)
	}

	wInfo := activity.GetInfo(ctx)
	buildDir := w.buildDir(in.Slug, sha)

	_, err := w.db.Queries().UpdateBuildStatus(ctx, sqlc.UpdateBuildStatusParams{
		MiddlewareSlug: in.Slug,
		Sha:            sha,
		Status:         middlewarev1.Build_RUNNING.String(),
		StatusDetail:   fmt.Sprintf("Running builder %v (attempt %d)", builder, wInfo.Attempt),
		OutputPath: sql.NullString{
			String: buildDir,
			Valid:  true,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to update build status: %w", err)
	}

	srcDir := w.srcDir(ctx, in.Slug, sha)
	logger.Info("Running builder", "Builder", builder, "SrcDir", srcDir)

	runWasmOut := filepath.Join(w.runDir(ctx, in.Slug, sha), "wasm.out")
	stdout, stderr, err := builder.Run(ctx, srcDir, runWasmOut, in.Params.BuildArgs...)
	if err != nil {
		return fmt.Errorf("builder failed: %w", err)
	}

	// Copy stdout and stderr to log.
	go func() {
		if err := copyLogs(buildDir, "stdout", stdout, int(wInfo.Attempt)); err != nil {
			logger.Error("Failed to copy stdout", "Error", err)
		}
	}()
	go func() {
		if err := copyLogs(buildDir, "stderr", stderr, int(wInfo.Attempt)); err != nil {
			logger.Error("Failed to copy stderr", "Error", err)
		}
	}()

	if err := builder.Wait(ctx); err != nil {
		return fmt.Errorf("failed waiting on builder: %w", err)
	}

	if _, err := os.Stat(runWasmOut); err != nil {
		return fmt.Errorf("wasm output not found: %w", err)
	}

	logger.Info("BuildActivity done", "Slug", in.Slug)

	return nil
}

func (w *IngestWorker) UploadWasmOutputActivity(ctx context.Context, slug, sha string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("UploadWasmFileActivity", "Slug", slug)

	if !w.remoteStorage {
		if err := os.Link(filepath.Join(w.runDir(ctx, slug, sha), "wasm.out"), w.wasmOut(ctx, slug, sha)); err != nil {
			return fmt.Errorf("failed to link wasm file: %w", err)
		}
		return nil
	} else {
		return errors.New("remote storage not implemented")
	}

	logger.Info("UploadWasmFileActivity done", "Slug", slug)

	return nil
}

func (w *IngestWorker) FinalizeActivity(ctx context.Context, in *MiddlewareIngestParams, r *MiddlewareIngestResult) error {
	logger := tlog.With(activity.GetLogger(ctx), "Slug", in.Slug, "SHA", r.SHA)
	logger.Info("finalizing build")

	status := middlewarev1.Build_READY
	statusDetail := "Ready"
	if r.Err != "" {
		status = middlewarev1.Build_ERRORED
		statusDetail = fmt.Sprintf("Failed: %v", r.Err)
	}

	if r.SHA != "" { // May not have a SHA if prepare failed.
		_, err := w.db.Queries().UpdateBuildStatus(ctx, sqlc.UpdateBuildStatusParams{
			MiddlewareSlug: in.Slug,
			Sha:            r.SHA,
			Status:         status.String(),
			StatusDetail:   statusDetail,
			OutputPath: sql.NullString{
				String: w.buildDir(in.Slug, r.SHA),
				Valid:  true,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to update build status: %w", err)
		}
	}

	m, err := w.db.Queries().GetMiddlewareBySlug(ctx, in.Slug)
	if err != nil {
		return fmt.Errorf("failed to get middleware: %w", err)
	}

	mStatus := middlewarev1.Middleware_READY
	mStatusDetail := fmt.Sprintf("Ready: %s")
	if r.Err != "" {
		// Initial pending is the only status that can go to errored.
		if m.Status == middlewarev1.Middleware_PENDING.String() {
			mStatus = middlewarev1.Middleware_ERRORED
			mStatusDetail = fmt.Sprintf("Errored: %v", r.Err)
		}
	}

	_, err = w.db.Queries().UpdateMiddlewareStatus(ctx, sqlc.UpdateMiddlewareStatusParams{
		Slug:   in.Slug,
		Status: mStatus.String(),
		StatusDetail: sql.NullString{
			String: mStatusDetail,
			Valid:  true,
		},
		LiveBuildSha: sql.NullString{
			String: r.SHA,
			Valid:  r.SHA != "",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to update middleware status: %w", err)
	}

	if in.Params.Type == middlewarev1.MiddlewareIngestParams_DIRECT {
		if err := w.watcher.Add(in.Slug, in.Params.WatchDir); err != nil {
			return fmt.Errorf("failed to add middleware to watcher: %w", err)
		}
	}

	logger.Info("build finalized")

	return nil
}
