// Package sl provides utilities for handling multi-handler logging using the slog package.
//
// The MultiHandler allows you to combine multiple slog handlers into a single handler.
// This is useful when you want to send log records to different outputs with varying
// configurations or levels, such as logging to both a file and the console.
//
// Example usage:
//
//	package main
//
//	import (
//	    "log/slog"
//	    "os"
//	    "github.com/repooooo/go-utils/sl"
//	)
//
//	// Global log variable that will use the MultiHandler
//	var log slog.Logger
//
//	func main() {
//	    // Create a log file to store JSON logs
//	    logFile, err := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
//	    if err != nil {
//	        panic(err)
//	    }
//
//	    // Initialize the global log variable with the MultiHandler
//	    log = slog.New(sl.NewMultiHandler(
//	        slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
//	        slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
//	    ))
//
//	    // Log some messages
//	    log.Debug("This is a debug message")
//	    log.Info("This is an info message")
//	}
//
// The example creates a logger that logs debug and info level messages to both the console
// (in text format) and a JSON log file.
package sl

import (
	"context"
	"log/slog"
)

// MultiHandler is a custom slog handler that allows multiple handlers to be combined.
// It sends log records to all registered handlers.
type MultiHandler struct {
	// handlers is a slice of slog.Handler instances that will handle log records.
	handlers []slog.Handler
}

// NewMultiHandler creates a new MultiHandler with the provided handlers.
// It accepts multiple slog.Handler instances and combines them into one handler.
func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

// Handle processes a log record by passing it to all handlers in the MultiHandler.
// It returns the first error encountered while processing the log record, or nil if all handlers succeed.
func (m *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, handler := range m.handlers {
		// Pass the record to each handler
		if err := handler.Handle(ctx, r); err != nil {
			return err // Return the first error encountered
		}
	}
	return nil
}

// Enabled checks if at least one handler in the MultiHandler is enabled for the specified log level.
// It returns true if any handler allows the log level, false otherwise.
func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range m.handlers {
		if handler.Enabled(ctx, level) {
			return true // At least one handler is enabled for the log level
		}
	}
	return false // No handlers are enabled for the log level
}

// WithAttrs adds additional attributes to each handler in the MultiHandler.
// It returns a new MultiHandler with the updated attributes for each handler.
func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var newHandlers []slog.Handler
	// Add the new attributes to each handler
	for _, handler := range m.handlers {
		newHandlers = append(newHandlers, handler.WithAttrs(attrs))
	}
	return &MultiHandler{handlers: newHandlers}
}

// WithGroup adds a group name to each handler in the MultiHandler.
// It returns a new MultiHandler with the updated group name for each handler.
func (m *MultiHandler) WithGroup(name string) slog.Handler {
	var newHandlers []slog.Handler
	// Add the new group name to each handler
	for _, handler := range m.handlers {
		newHandlers = append(newHandlers, handler.WithGroup(name))
	}
	return &MultiHandler{handlers: newHandlers}
}
