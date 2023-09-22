package slogdiscard

import (
	"context"

	"golang.org/x/exp/slog"
)

// NewDiscardLogger create new DiscardLogger
func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

// Type for DiscartLogger
type DiscardHandler struct{}

// New handler for DiscardLogger
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// Stub
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// Stub
func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// Stub
func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}

// Stub
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
