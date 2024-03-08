package slog2

import (
	"context"
	"log/slog"

	"github.com/go-logr/logr"
)

type slogContextKey struct{}

func DefaultOfContext(ctx context.Context) *slog.Logger {
	if v := ctx.Value(slogContextKey{}); v != nil {
		return v.(*slog.Logger)
	} else {
		return logr.FromContextAsSlogLogger(ctx)
	}
}

func FromContextOrDefault(ctx context.Context) *slog.Logger {
	if v := DefaultOfContext(ctx); v != nil {
		return v
	}
	return slog.Default()
}

func NewContext(parent context.Context, l *slog.Logger) context.Context {
	return context.WithValue(parent, slogContextKey{}, l)
}
