package envoy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/apoxy-dev/proximal/core/log"
	"github.com/google/uuid"
)

const (
	githubURL = "github.com/envoyproxy/envoy/releases/download"
)

var (
	goArchToPlatform = map[string]string{
		"amd64": "x86_64",
		"arm64": "aarch_64",
	}
)

type Release struct {
	Version string
	Sha     string
}

func (r *Release) String() string {
	if r.Sha == "" {
		return r.Version
	}
	return fmt.Sprintf("%s@sha256:%s", r.Version, r.Sha)
}

func (r *Release) DownloadBinaryFromGitHub(ctx context.Context) (io.ReadCloser, error) {
	downloadURL := filepath.Join(
		githubURL,
		r.Version,
		fmt.Sprintf("envoy-%s-%s-%s", r.Version[1:], runtime.GOOS, goArchToPlatform[runtime.GOARCH]),
	)

	log.Infof("downloading envoy %s from https://%s", r, downloadURL)

	resp, err := http.Get("https://" + downloadURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download envoy: %w", err)
	}
	return resp.Body, nil
}

// Runtime vendors the Envoy binary and runs it.
type Runtime struct {
	EnvoyPath           string
	BootstrapConfigPath string
	BootstrapConfigYAML string
	Release             *Release
	// Args are additional arguments to pass to Envoy.
	Args []string

	cmd *exec.Cmd
}

func (r *Runtime) run(ctx context.Context) error {
	id := uuid.New().String()
	configYAML := fmt.Sprintf(`node: { id: "%s", cluster: "proximal" }`, id)
	if r.BootstrapConfigYAML != "" {
		configYAML = r.BootstrapConfigYAML
	}
	log.Infof("envoy YAML config: %s", configYAML)

	args := []string{
		"--config-yaml", configYAML,
	}

	if r.BootstrapConfigPath != "" {
		args = append(args, "-c", r.BootstrapConfigPath)
	}

	args = append(args, r.Args...)

	r.cmd = exec.CommandContext(ctx, r.envoyPath(), args...)
	r.cmd.Stdout = os.Stdout
	r.cmd.Stderr = os.Stderr

	if err := r.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start envoy: %w", err)
	}

	// Restart envoy if it exits.
	if err := r.cmd.Wait(); err != nil {
		return fmt.Errorf("envoy exited with error: %w", err)
	}

	return nil
}

// envoyPath returns the path to the Envoy binary. If EnvoyPath is set, it will
// be used. Otherwise, the binary will be downloaded and cached in
// ~/.proximal/envoy for each release.
func (r *Runtime) envoyPath() string {
	if r.EnvoyPath != "" {
		return r.EnvoyPath
	}
	return filepath.Join(os.Getenv("HOME"), ".proximal", "envoy", r.Release.String(), "bin", "envoy")
}

// vendorEnvoyIfNotExists vendors the Envoy binary for the release if it does
// not exist.
func (r *Runtime) vendorEnvoyIfNotExists(ctx context.Context) error {
	if _, err := os.Stat(r.envoyPath()); err == nil {
		return nil
	}

	// Download the Envoy binary for the release.
	bin, err := r.Release.DownloadBinaryFromGitHub(ctx)
	if err != nil {
		return fmt.Errorf("failed to download envoy: %w", err)
	}
	defer bin.Close()

	// Extract the Envoy binary.
	if err := os.MkdirAll(filepath.Dir(r.envoyPath()), 0755); err != nil {
		return fmt.Errorf("failed to create envoy directory: %w", err)
	}
	w, err := os.OpenFile(r.envoyPath(), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("failed to open envoy: %w", err)
	}
	defer w.Close()
	if _, err := io.Copy(w, bin); err != nil {
		return fmt.Errorf("failed to copy envoy: %w", err)
	}
	if err := os.Chmod(r.envoyPath(), 0755); err != nil {
		return fmt.Errorf("failed to chmod envoy: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(r.BootstrapConfigPath), 0755); err != nil {
		return fmt.Errorf("failed to create envoy directory: %w", err)
	}

	return nil
}

// Run runs the Envoy binary.
func (r *Runtime) Run(ctx context.Context) error {
	log.Infof("running envoy %s", r.Release)

	if err := r.vendorEnvoyIfNotExists(ctx); err != nil {
		return fmt.Errorf("failed to vendor envoy: %w", err)
	}

	// Run the Envoy binary.
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		exitCh := make(chan struct{})
		go func() {
			if err := r.run(ctx); err != nil {
				log.Errorf("envoy exited with error: %v", err)
			}
			close(exitCh)
		}()

		select {
		case <-ctx.Done():
			return nil
		case <-exitCh:
		}
	}

	return nil
}

// Stop stops the Envoy process.
func (r *Runtime) Stop() error {
	if r.cmd == nil {
		return nil
	}
	return r.cmd.Process.Kill()
}
