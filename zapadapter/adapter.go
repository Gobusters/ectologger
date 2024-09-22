package zapadapter

import (
	"github.com/Gobusters/ectologger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Helper function to convert map[string]interface{} to []zap.Field
func fieldsToZapFields(fields map[string]interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return zapFields
}

// GetZapLogFunc returns a log function that logs to the provided zap logger
// before is a function that is called before the log message is logged.
// It can be used to modify the log message or add additional fields to it.
func GetZapLogFunc(zapLogger *zap.Logger, before func(msg ectologger.EctoLogMessage) ectologger.EctoLogMessage) ectologger.EctoLogFunc {
	return func(msg ectologger.EctoLogMessage) {
		if before != nil {
			msg = before(msg)
		}

		zapFields := fieldsToZapFields(msg.Fields)
		level, err := zapcore.ParseLevel(msg.Level)
		if err != nil {
			level = zapcore.InfoLevel // Default to Info level if parsing fails
		}

		if msg.Err != nil {
			zapFields = append(zapFields, zap.Error(msg.Err))
		}

		zapLogger.Log(level, msg.Message, zapFields...)
	}
}

// NewZapEctoLogger returns a new EctoLogger that logs to the provided zap logger
// before is an optional function that is called before the log message is logged.
// It can be used to modify the log message or add additional fields to it.
func NewZapEctoLogger(zapLogger *zap.Logger, before func(msg ectologger.EctoLogMessage) ectologger.EctoLogMessage) ectologger.Logger {
	return ectologger.NewEctoLogger(GetZapLogFunc(zapLogger, before))
}
