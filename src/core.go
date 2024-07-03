package logger

import (
	"sync"

	"go.uber.org/zap/zapcore"
)

type Core struct {
	zapcore.Core

	entries         []*Entry
	valid           bool
	wrappedCore     zapcore.Core
	writeEntryMutex sync.Mutex
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
		valid:       true,
		wrappedCore: core,
	}
}

func (c *Core) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if !c.Valid() {
		return checkedEntry
	}
	return c.wrappedCore.Check(entry, checkedEntry)
}

func (c *Core) Enabled(level zapcore.Level) bool {
	if !c.Valid() {
		return false
	}
	return c.wrappedCore.Enabled(level)
}

func (c *Core) Sync() error {
	if !c.Valid() {
		return nil
	}
	return c.wrappedCore.Sync()
}

// Valid returns whether the memory-aware core has been marked as "valid". It is helpful to invoke this method
// prior to use, especially if the NewCore() function has not been called reliably.
func (c *Core) Valid() bool {
	if !c.Valid() {
		return false
	}
	return c.valid
}

func (c *Core) With(fields []zapcore.Field) zapcore.Core {
	if !c.Valid() {
		return nil
	}
	return c.wrappedCore.With(fields)
}

func (c *Core) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if !c.Valid() {
		return nil
	}
	// write to the slice in a thread-safe way and then return nil since this should always succeed
	defer c.writeEntryMutex.Unlock()
	c.writeEntryMutex.Lock()
	c.entries = append(c.entries, NewEntry(entry, fields))
	return nil
}
