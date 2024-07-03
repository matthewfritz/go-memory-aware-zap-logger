package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	memoryAwareCore *Core
	valid           bool
	wrappedLogger   *zap.Logger
}

// NewLogger takes a zap.Logger pointer and returns a wrapped memory-aware logger pointer.
func NewLogger(logger *zap.Logger) *Logger {
	if logger == nil {
		return &Logger{
			wrappedLogger: zap.NewNop(),
		}
	}
	memoryAwareCore := NewCore(logger.Core())
	wrapOption := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, memoryAwareCore)
	})
	return &Logger{
		memoryAwareCore: memoryAwareCore,
		valid:           true,
		wrappedLogger:   logger.WithOptions(wrapOption),
	}
}

// Valid returns whether the memory-aware logger has been marked as "valid". It is helpful to invoke this method
// prior to use, especially if the NewLogger() function has not been called reliably.
func (l *Logger) Valid() bool {
	if l == nil || l.wrappedLogger == nil {
		return false
	}
	return l.valid
}

// WrappedLogger returns the underlying zap.Logger pointer so you can continue chaining log options.
func (l *Logger) WrappedLogger() *zap.Logger {
	if !l.Valid() {
		return zap.NewNop()
	}
	return l.wrappedLogger
}
