package memorylogger

import "go.uber.org/zap/zapcore"

type Entry struct {
	entry  zapcore.Entry
	fields []zapcore.Field
}

// NewEntry takes a zapcore.Entry instance and a slice of zapcore.Field instances and returns a new
// Entry pointer for memory-aware logging.
func NewEntry(entry zapcore.Entry, fields []zapcore.Field) *Entry {
	return &Entry{
		entry,
		fields,
	}
}
