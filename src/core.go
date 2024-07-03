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
	if core == nil {
		return &Core{
			entries:     []*Entry{},
			wrappedCore: zapcore.NewNopCore(),
		}
	}
	return &Core{
		entries:     []*Entry{},
		wrappedCore: core,
	}
}

func (c *Core) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c == nil {
		return checkedEntry
	}
	return c.wrappedCore.Check(entry, checkedEntry)
}

func (c *Core) Enabled(level zapcore.Level) bool {
	if c == nil {
		return false
	}
	return c.wrappedCore.Enabled(level)
}

func (c *Core) Sync() error {
	if c == nil {
		return nil
	}
	return c.wrappedCore.Sync()
}

func (c *Core) With(fields []zapcore.Field) zapcore.Core {
	if c == nil {
		return nil
	}
	return c.wrappedCore.With(fields)
}

func (c *Core) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if c == nil {
		return nil
	}
	// write to the slice and then return nil since this should always succeed
	c.entries = append(c.entries, NewEntry(entry, fields))
	return nil
}
