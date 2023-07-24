package gobuild

import (
	"context"
	"io"
	"os/exec"
)

type Builder struct {
	cmd *exec.Cmd
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) String() string {
	return "tinygo"
}

func (b *Builder) Run(ctx context.Context, cwd, output string, args ...string) (stdout, stderr io.Reader, err error) {
	defaultArgs := []string{
		"build",
		"-o", output,
		"-target", "wasi",
		"-scheduler", "none",
	}
	b.cmd = exec.CommandContext(
		ctx,
		"tinygo", append(defaultArgs, args...)...,
	)
	b.cmd.Dir = cwd
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
	return b.cmd.Wait()
}
