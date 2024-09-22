package ectologger

import (
	"context"
)

// Logger is an interface for logging operations.
// It provides methods for logging at different levels (Debug, Info, Warn, Error, Fatal)
// and allows for adding contextual information through fields and context.
//
// While this interface is implemented by ectologger, Go conventions recommend
// using your own logger interface in your code.
type Logger interface {
	// WithFields returns a new Logger with the given fields added to the logging context.
	WithFields(fields map[string]interface{}) Logger

	// WithField returns a new Logger with the given key-value pair added to the logging context.
	WithField(key string, value interface{}) Logger

	// WithContext returns a new Logger with the given context added to the logging context.
	WithContext(ctx context.Context) Logger

	// WithError returns a new Logger with the given error added to the logging context.
	WithError(err error) Logger

	// Debug logs a message at the Debug level.
	Debug(msg string)

	// Debugf logs a formatted message at the Debug level.
	Debugf(format string, args ...any)

	// DebugContext logs a message at the Debug level with the given context.
	DebugContext(ctx context.Context, msg string)

	// DebugContextf logs a formatted message at the Debug level with the given context.
	DebugContextf(ctx context.Context, format string, args ...any)

	// Info logs a message at the Info level.
	Info(msg string)

	// Infof logs a formatted message at the Info level.
	Infof(format string, args ...any)

	// InfoContext logs a message at the Info level with the given context.
	InfoContext(ctx context.Context, msg string)

	// InfoContextf logs a formatted message at the Info level with the given context.
	InfoContextf(ctx context.Context, format string, args ...any)

	// Warn logs a message at the Warn level.
	Warn(msg string)

	// Warnf logs a formatted message at the Warn level.
	Warnf(format string, args ...any)

	// WarnContext logs a message at the Warn level with the given context.
	WarnContext(ctx context.Context, msg string)

	// WarnContextf logs a formatted message at the Warn level with the given context.
	WarnContextf(ctx context.Context, format string, args ...any)

	// Error logs a message at the Error level.
	Error(msg string)

	// Errorf logs a formatted message at the Error level.
	Errorf(format string, args ...any)

	// ErrorContext logs a message at the Error level with the given context.
	ErrorContext(ctx context.Context, msg string)

	// ErrorContextf logs a formatted message at the Error level with the given context.
	ErrorContextf(ctx context.Context, format string, args ...any)

	// Fatal logs a message at the Fatal level
	Fatal(msg string)

	// Fatalf logs a formatted message at the Fatal level
	Fatalf(format string, args ...any)

	// FatalContext logs a message at the Fatal level with the given context
	FatalContext(ctx context.Context, msg string)

	// FatalContextf logs a formatted message at the Fatal level with the given context
	FatalContextf(ctx context.Context, format string, args ...any)
}
