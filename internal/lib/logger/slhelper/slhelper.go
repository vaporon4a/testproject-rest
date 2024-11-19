package slhelper

import "log/slog"

// Err returns an slog.Attr for the given error.
//
// The attribute has the key "error" and the value is a string representation
// of the error.
//
// This can be used when you want to log an error, but not as a side effect of
// the log statement failing, e.g. slog.Error("failed to foo", Err(err))
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
