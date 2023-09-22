package sl

import "golang.org/x/exp/slog"

// Converts regular error to error for log/slog
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
