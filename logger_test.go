package ectologger

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEctoLogger(t *testing.T) {
	logFunc := func(msg EctoLogMessage) {
		// do nothing
	}
	logger := NewEctoLogger(logFunc)

	assert.NotNil(t, logger)
	assert.IsType(t, &EctoLogger{}, logger)
}

func TestNewDefaultEctoLogger(t *testing.T) {
	logger := NewDefaultEctoLogger()

	assert.NotNil(t, logger)
	assert.IsType(t, &EctoLogger{}, logger)
}

func TestEctoLoggerWithFields(t *testing.T) {
	originalLogger := NewDefaultEctoLogger()
	fields := map[string]interface{}{"key": "value"}

	subLogger := originalLogger.WithFields(fields)

	assert.NotNil(t, subLogger)
	assert.IsType(t, &ectoSubLogger{}, subLogger)
}

func TestEctoLoggerWithField(t *testing.T) {
	originalLogger := NewDefaultEctoLogger()

	subLogger := originalLogger.WithField("key", "value")

	assert.NotNil(t, subLogger)
	assert.IsType(t, &ectoSubLogger{}, subLogger)
}

func TestEctoLoggerWithContext(t *testing.T) {
	originalLogger := NewDefaultEctoLogger()
	ctx := context.Background()

	subLogger := originalLogger.WithContext(ctx)

	assert.NotNil(t, subLogger)
	assert.IsType(t, &ectoSubLogger{}, subLogger)
}

func TestEctoLoggerWithError(t *testing.T) {
	originalLogger := NewDefaultEctoLogger()
	err := errors.New("test error")

	subLogger := originalLogger.WithError(err)

	assert.NotNil(t, subLogger)
	assert.IsType(t, &ectoSubLogger{}, subLogger)
}

func TestEctoLoggerLogMethods(t *testing.T) {
	testCases := []struct {
		name     string
		logLevel string
		logFunc  func(l Logger, msg string)
	}{
		{"Debug", "debug", func(l Logger, msg string) { l.Debug(msg) }},
		{"Info", "info", func(l Logger, msg string) { l.Info(msg) }},
		{"Warn", "warn", func(l Logger, msg string) { l.Warn(msg) }},
		{"Error", "error", func(l Logger, msg string) { l.Error(msg) }},
		{"Fatal", "fatal", func(l Logger, msg string) { l.Fatal(msg) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var capturedMsg EctoLogMessage
			logFunc := func(msg EctoLogMessage) {
				capturedMsg = msg
			}

			logger := NewEctoLogger(logFunc)
			tc.logFunc(logger, "test message")

			assert.Equal(t, tc.logLevel, capturedMsg.Level)
			assert.Equal(t, "test message", capturedMsg.Message)
		})
	}
}

func TestEctoSubLoggerLogMethods(t *testing.T) {
	testCases := []struct {
		name     string
		logLevel string
		logFunc  func(l Logger, msg string)
	}{
		{"Debug", "debug", func(l Logger, msg string) { l.Debug(msg) }},
		{"Info", "info", func(l Logger, msg string) { l.Info(msg) }},
		{"Warn", "warn", func(l Logger, msg string) { l.Warn(msg) }},
		{"Error", "error", func(l Logger, msg string) { l.Error(msg) }},
		{"Fatal", "fatal", func(l Logger, msg string) { l.Fatal(msg) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var capturedMsg EctoLogMessage
			logFunc := func(msg EctoLogMessage) {
				capturedMsg = msg
			}

			type contextKey string

			const testContextKey contextKey = "testKey"

			fields := map[string]interface{}{"key": "value"}
			ctx := context.WithValue(context.Background(), testContextKey, "testValue")
			err := errors.New("test error")

			logger := NewEctoLogger(logFunc).
				WithFields(fields).
				WithContext(ctx).
				WithError(err)

			tc.logFunc(logger, "test message")

			assert.Equal(t, tc.logLevel, capturedMsg.Level)
			assert.Equal(t, "test message", capturedMsg.Message)
			assert.Equal(t, fields, capturedMsg.Fields)
			assert.Equal(t, ctx, capturedMsg.Ctx)
			assert.Equal(t, err, capturedMsg.Err)
		})
	}
}

func TestDefaultEctoLogFunc(t *testing.T) {
	msg := EctoLogMessage{
		Level:   "info",
		Message: "test message",
		Fields:  map[string]interface{}{"key": "value"},
		Ctx:     context.Background(),
		Err:     errors.New("test error"),
	}

	// Capture the output of log.Print
	var logOutput string
	log.SetOutput(writerFunc(func(p []byte) (int, error) {
		logOutput = string(p)
		return len(p), nil
	}))

	DefaultEctoLogFunc(msg)

	// Parse the JSON output
	var parsedOutput map[string]interface{}
	err := json.Unmarshal([]byte(logOutput), &parsedOutput)
	require.NoError(t, err)

	// Assert the parsed output
	assert.Equal(t, "info", parsedOutput["level"])
	assert.Equal(t, "test message", parsedOutput["message"])
	assert.Equal(t, "test error", parsedOutput["err"])
	assert.Equal(t, "value", parsedOutput["key"])
	assert.NotEmpty(t, parsedOutput["time"])
}

// writerFunc is a helper type to capture log output
type writerFunc func(p []byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) {
	return f(p)
}
