package gobuild

import (
	"context"
	"io"
	"os/exec"
)

type Builder struct {
	modCmd, cmd *exec.Cmd
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) String() string {
	return "tinygo"
}

func (b *Builder) Run(ctx context.Context, cwd, output string, args ...string) (stdout, stderr io.Reader, err error) {
	modArgs := []string{
		"mod",
		"download",
	}
	b.modCmd = exec.CommandContext(
		ctx,
		"go", modArgs...,
	)
	b.modCmd.Dir = cwd
	modOut, err := b.modCmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	modErr, err := b.modCmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	if err := b.modCmd.Start(); err != nil {
		return nil, nil, err
	}

	buildArgs := []string{
		"build",
		"-o", output,
		"-target", "wasi",
		"-scheduler", "none",
	}
	b.cmd = exec.CommandContext(
		ctx,
		"tinygo", append(buildArgs, args...)...,
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

	return io.MultiReader(modOut, stdout), io.MultiReader(modErr, stderr), nil
}

func (b *Builder) Wait(ctx context.Context) error {
	if err := b.modCmd.Wait(); err != nil {
		return err
	}
	return b.cmd.Wait()
}
