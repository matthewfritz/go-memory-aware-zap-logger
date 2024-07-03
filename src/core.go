package logger

import (
	"go.uber.org/zap/zapcore"
)

type Core struct {
	zapcore.Core

	entries     []*Entry
	wrappedCore zapcore.Core
}

// NewCore takes a zapcore.Core implementation and returns a new memory-aware logging Core pointer.
func NewCore(core zapcore.Core) *Core {
	return &Core{
		entries:     []*Entry{},
		wrappedCore: core,
	}
}

func (c *Core) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return c.wrappedCore.Check(entry, checkedEntry)
}

func (c *Core) Enabled(level zapcore.Level) bool {
	return c.wrappedCore.Enabled(level)
}

func (c *Core) Sync() error {
	return c.wrappedCore.Sync()
}

func (c *Core) With(fields []zapcore.Field) zapcore.Core {
	return c.wrappedCore.With(fields)
}

func (c *Core) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	// write to the slice and then return nil since this should always succeed
	c.entries = append(c.entries, NewEntry(entry, fields))
	return nil
}
