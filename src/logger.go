package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	memoryAwareCore *Core
	wrappedLogger   *zap.Logger
}

// NewLogger takes a zap.Logger pointer and returns a wrapped memory-aware logger pointer.
func NewLogger(logger *zap.Logger) *Logger {
	if logger == nil {
		return nil
	}
	memoryAwareCore := NewCore(logger.Core())
	wrapOption := zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, memoryAwareCore)
	})
	return &Logger{
		wrappedLogger: logger.WithOptions(wrapOption),
	}
}

// WrappedLogger returns the underlying zap.Logger pointer so you can continue chaining log options.
func (l *Logger) WrappedLogger() *zap.Logger {
	return l.wrappedLogger
}
