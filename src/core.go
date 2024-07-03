package memorylogger

import (
	"go.uber.org/zap/zapcore"
)

type Core struct {
	zapcore.Core

	entries     []zapcore.Entry
	wrappedCore zapcore.Core
}

// NewCore takes a zapcore.Core implementation and returns a new memory-aware logging Core pointer.
func NewCore(core zapcore.Core) *Core {
	return &Core{
		entries:     []zapcore.Entry{},
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
	// TODO: write to the memory buffer but for now just do a pass-through to check doubled log lines
	return c.wrappedCore.Write(entry, fields)
}
