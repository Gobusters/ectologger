package ectologger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Gobusters/ectolinq"
)

// EctoLogMessage represents a log message with its associated metadata.
type EctoLogMessage struct {
	Level   string                 // The log level of the message. One of: debug, info, warn, error, fatal
	Message string                 // The log message
	Fields  map[string]interface{} // Fields to add to the log message
	Ctx     context.Context        // The context of the log message
	Err     error                  // The error to add to the log message
}

// EctoLogFunc is a function type that defines how a log message should be processed.
type EctoLogFunc func(msg EctoLogMessage)

// EctoLogger is the main logger struct that implements the Logger interface.
type EctoLogger struct {
	logFunc EctoLogFunc
}

// NewEctoLogger creates a new EctoLogger with the given log function.
func NewEctoLogger(logFunc EctoLogFunc) Logger {
	return &EctoLogger{logFunc: logFunc}
}

// DefaultEctoLogFunc is the default log function.
// It marshals the log message to JSON and writes it to stdout.
func DefaultEctoLogFunc(msg EctoLogMessage) {
	jsonMsg := make(map[string]interface{}, len(msg.Fields)+4) // Pre-allocate map with estimated size
	jsonMsg["level"] = msg.Level
	jsonMsg["message"] = msg.Message
	jsonMsg["err"] = msg.Err
	jsonMsg["time"] = time.Now().Format(time.RFC3339)

	jsonMsg = ectolinq.Merge(jsonMsg, msg.Fields)

	json, err := json.Marshal(jsonMsg)
	if err != nil {
		log.Printf("Error marshalling log message to JSON: %v", err)
		return
	}

	log.Print(string(json)) // Avoid unnecessary string formatting
}

// NewDefaultEctoLogger returns a new EctoLogger that logs to the default logger
func NewDefaultEctoLogger() Logger {
	return NewEctoLogger(DefaultEctoLogFunc)
}

// WithFields returns a new Logger with the given fields added to the logging context.
func (l *EctoLogger) WithFields(fields map[string]interface{}) Logger {
	return &ectoSubLogger{logFunc: l.logFunc, fields: fields}
}

// WithField returns a new Logger with the given key-value pair added to the logging context.
func (l *EctoLogger) WithField(key string, value interface{}) Logger {
	return &ectoSubLogger{logFunc: l.logFunc, fields: map[string]interface{}{key: value}}
}

// WithContext returns a new Logger with the given context added to the logging context.
func (l *EctoLogger) WithContext(ctx context.Context) Logger {
	return &ectoSubLogger{logFunc: l.logFunc, fields: map[string]interface{}{}, ctx: ctx}
}

// WithError returns a new Logger with the given error added to the logging context.
func (l *EctoLogger) WithError(err error) Logger {
	return &ectoSubLogger{logFunc: l.logFunc, fields: map[string]interface{}{}, err: err}
}

// Debug logs a message at the Debug level.
func (l *EctoLogger) Debug(msg string) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) Debugf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) DebugContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}
func (l *EctoLogger) DebugContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}

// Info logs a message at the Info level.
func (l *EctoLogger) Info(msg string) {
	l.logFunc(EctoLogMessage{Level: "info", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) Infof(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "info", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) InfoContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "info", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}
func (l *EctoLogger) InfoContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "info", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}

// Warn logs a message at the Warn level.
func (l *EctoLogger) Warn(msg string) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) Warnf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) WarnContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}
func (l *EctoLogger) WarnContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}

// Error logs a message at the Error level.
func (l *EctoLogger) Error(msg string) {
	l.logFunc(EctoLogMessage{Level: "error", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) Errorf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "error", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) ErrorContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "error", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}
func (l *EctoLogger) ErrorContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "error", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}

// Fatal logs a message at the Fatal level.
func (l *EctoLogger) Fatal(msg string) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) Fatalf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: nil})
}
func (l *EctoLogger) FatalContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: msg, Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}
func (l *EctoLogger) FatalContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: fmt.Sprintf(format, args...), Fields: map[string]interface{}{}, Err: nil, Ctx: ctx})
}

// ectoSubLogger is an internal type that represents a logger with additional context.
type ectoSubLogger struct {
	logFunc EctoLogFunc
	fields  map[string]interface{}
	err     error
	ctx     context.Context
}

// WithFields returns a new Logger with the given fields added to the logging context.
func (l *ectoSubLogger) WithFields(fields map[string]interface{}) Logger {
	l.fields = ectolinq.Merge(l.fields, fields)
	return l
}

// WithField returns a new Logger with the given key-value pair added to the logging context.
func (l *ectoSubLogger) WithField(key string, value interface{}) Logger {
	l.fields = ectolinq.Merge(l.fields, map[string]interface{}{key: value})
	return l
}

// WithContext returns a new Logger with the given context added to the logging context.
func (l *ectoSubLogger) WithContext(ctx context.Context) Logger {
	l.ctx = ctx
	return l
}

// WithError returns a new Logger with the given error added to the logging context.
func (l *ectoSubLogger) WithError(err error) Logger {
	l.err = err
	return l
}

// Debug logs a message at the Debug level.
func (l *ectoSubLogger) Debug(msg string) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: msg, Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) Debugf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) DebugContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: msg, Fields: l.fields, Err: l.err, Ctx: ctx})
}
func (l *ectoSubLogger) DebugContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "debug", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: ctx})
}

// Info logs a message at the Info level.
func (l *ectoSubLogger) Info(msg string) {
	l.logFunc(EctoLogMessage{Level: "info", Message: msg, Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) Infof(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "info", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) InfoContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "info", Message: msg, Fields: l.fields, Err: l.err, Ctx: ctx})
}
func (l *ectoSubLogger) InfoContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "info", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: ctx})
}

// Warn logs a message at the Warn level.
func (l *ectoSubLogger) Warn(msg string) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: msg, Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) Warnf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) WarnContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: msg, Fields: l.fields, Err: l.err, Ctx: ctx})
}
func (l *ectoSubLogger) WarnContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "warn", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: ctx})
}

// Error logs a message at the Error level.
func (l *ectoSubLogger) Error(msg string) {
	l.logFunc(EctoLogMessage{Level: "error", Message: msg, Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) Errorf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "error", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) ErrorContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "error", Message: msg, Fields: l.fields, Err: l.err, Ctx: ctx})
}
func (l *ectoSubLogger) ErrorContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "error", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: ctx})
}

// Fatal logs a message at the Fatal level.
func (l *ectoSubLogger) Fatal(msg string) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: msg, Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) Fatalf(format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: l.ctx})
}
func (l *ectoSubLogger) FatalContext(ctx context.Context, msg string) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: msg, Fields: l.fields, Err: l.err, Ctx: ctx})
}
func (l *ectoSubLogger) FatalContextf(ctx context.Context, format string, args ...any) {
	l.logFunc(EctoLogMessage{Level: "fatal", Message: fmt.Sprintf(format, args...), Fields: l.fields, Err: l.err, Ctx: ctx})
}
