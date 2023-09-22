package slogdiscard

import (
	"context"

	"golang.org/x/exp/slog"
)

// NewDiscardLogger create new DiscardLogger
func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

// DiscardHandler is type for handle DiscartLogger
type DiscardHandler struct{}

// NewDiscardHandler returns handler for DiscardLogger
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// Handle return nil
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs return handler of DiscardHandler
func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup return handler of DiscardHandler
func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}

// Enabled return false
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
