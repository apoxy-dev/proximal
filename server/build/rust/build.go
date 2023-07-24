package rustbuild

import (
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/apoxy-dev/proximal/core/log"
)

type Builder struct {
	cmd    *exec.Cmd
	output string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) String() string {
	return "rustc (wasm32-wasi)"
}

func (b *Builder) Run(ctx context.Context, cwd, output string, args ...string) (stdout, stderr io.Reader, err error) {
	defaultArgs := []string{
		"build",
		"--target", "wasm32-wasi",
		"--target-dir", filepath.Join(cwd, "target"),
		"--release",
	}
	b.cmd = exec.CommandContext(
		ctx,
		"cargo", append(defaultArgs, args...)...,
	)
	b.cmd.Dir = cwd
	b.output = output
	stdout, err = b.cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err = b.cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	if err := b.cmd.Start(); err != nil {
		return nil, nil, err
	}

	return stdout, stderr, nil
}

func (b *Builder) Wait(ctx context.Context) error {
	if err := b.cmd.Wait(); err != nil {
		return err
	}

	// Move the file from the cargo output directory to the provided output path.
	log.Infof("searching in: %v", filepath.Join(b.cmd.Dir, "target/wasm32-wasi/release"))
	files, err := os.ReadDir(filepath.Join(b.cmd.Dir, "target/wasm32-wasi/release"))
	if err != nil {
		return err
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".wasm" {
			log.Debugf(
				"found output wasm: %v -> %v",
				filepath.Join(b.cmd.Dir, "target/wasm32-wasi/release", f.Name()),
				b.output,
			)
			if err := os.Rename(filepath.Join(b.cmd.Dir, "target/wasm32-wasi/release", f.Name()), b.output); err != nil {
				return err
			}
			break
		}
	}

	return nil
}
