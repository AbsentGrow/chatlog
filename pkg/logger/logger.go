// Package logger provides a configurable logging system for IPC service.
// Logs are written to stderr to keep stdout clean for IPC communication.
package logger

import (
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
)

var (
	log  zerolog.Logger
	mu   sync.RWMutex
	once sync.Once
)

// Init initializes the logger with the specified level.
// Level can be: debug, info, warn, error
func Init(level string) {
	once.Do(func() {
		setLevel(level)
	})
}

// SetLevel dynamically changes the log level.
func SetLevel(level string) {
	mu.Lock()
	defer mu.Unlock()
	setLevel(level)
}

func setLevel(level string) {
	var lvl zerolog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = zerolog.DebugLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "warn", "warning":
		lvl = zerolog.WarnLevel
	case "error":
		lvl = zerolog.ErrorLevel
	default:
		lvl = zerolog.InfoLevel
	}

	// Output to stderr to keep stdout for IPC
	log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"}).
		With().
		Timestamp().
		Logger().
		Level(lvl)
}

// GetLogger returns the global logger instance.
func GetLogger() *zerolog.Logger {
	mu.RLock()
	defer mu.RUnlock()
	return &log
}

// Debug logs a debug message.
func Debug() *zerolog.Event {
	mu.RLock()
	defer mu.RUnlock()
	return log.Debug()
}

// Info logs an info message.
func Info() *zerolog.Event {
	mu.RLock()
	defer mu.RUnlock()
	return log.Info()
}

// Warn logs a warning message.
func Warn() *zerolog.Event {
	mu.RLock()
	defer mu.RUnlock()
	return log.Warn()
}

// Error logs an error message.
func Error() *zerolog.Event {
	mu.RLock()
	defer mu.RUnlock()
	return log.Error()
}
