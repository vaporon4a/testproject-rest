package sldiscard

import (
	"context"
	"log/slog"
)

// NewDiscardLogger creates a new slog.Logger instance with a DiscardHandler,
// which discards all log records. This logger can be used in scenarios where
// logging is not necessary, such as in tests or when performance is critical.
func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct{}

// NewDiscardHandler creates a new instance of DiscardHandler.
// This handler effectively ignores all log records, making it
// suitable for use in testing or scenarios where logging is not required.
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// Handle is a no-op method that discards the log record.
// It implements the slog.Handler interface by returning nil,
// indicating a successful operation without processing the record.
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs is a no-op method that simply returns the receiver.
// It implements the slog.HandlerWithAttrs interface, but since
// the handler discards all records, there is no need to store or
// process the attributes.
func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup is a no-op method that simply returns the receiver.
// It implements the slog.HandlerWithGroup interface, but since
// the handler discards all records, there is no need to store or
// process the group name.
func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}

// Enabled always returns false, indicating that no log levels are enabled.
// This method implements the slog.Handler interface, effectively disabling
// logging when using the DiscardHandler.
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
