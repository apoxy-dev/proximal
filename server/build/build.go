package build

import (
	"context"
	"io"
)

type Builder interface {
	Run(ctx context.Context, cwd, output string, args ...string) (stdout, stderr io.Reader, err error)
	Wait(ctx context.Context) error
	String() string
}
