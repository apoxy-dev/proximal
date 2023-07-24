// Package log provides logging routines based on slog package.
package log

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/exp/slog"
)

var logger *slog.Logger

func init() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove time.
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			a.Value = slog.StringValue(filepath.Base(a.Value.String()))
		}
		return a
	}
	logger = slog.New(slog.HandlerOptions{AddSource: true, ReplaceAttr: replace}.NewTextHandler(os.Stdout))
	slog.SetDefault(logger)
}

func logf(level slog.Level, format string, args ...any) {
	ctx := context.Background()
	if !logger.Enabled(ctx, level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip [Callers, logf, Infof]
	r := slog.NewRecord(time.Now(), level, fmt.Sprintf(format, args...), pcs[0])
	_ = logger.Handler().Handle(ctx, r)
}

// Debugf logs a debug message.
func Debugf(format string, args ...any) {
	level := slog.LevelDebug
	logf(level, format, args...)
}

// Infof logs an info message.
func Infof(format string, args ...any) {
	level := slog.LevelInfo
	logf(level, format, args...)
}

// Warnf logs a warning message.
func Warnf(format string, args ...any) {
	level := slog.LevelWarn
	logf(level, format, args...)
}

// Errorf logs an error message.
func Errorf(format string, args ...any) {
	level := slog.LevelError
	logf(level, format, args...)
}

// Fatalf logs a fatal message.
func Fatalf(format string, args ...any) {
	level := slog.LevelError
	logf(level, format, args...)
	os.Exit(1)
}
