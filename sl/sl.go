// Package sl github.com/repooooo/go-utils/sl/sl.go
package sl

import "log/slog"

// Err returns an attribute containing the error message for logging purposes.
// The returned attribute can be used with Go's `log/slog` package to log error information in a structured way.
// This function extracts the error message from the given error object and formats it as a slog attribute.
// The key of the attribute is "error", and the value is the string representation of the error message.
//
// Example usage:
//
//	err := fmt.Errorf("something went wrong")
//	log := slog.New(slog.NewTextHandler(os.Stdout))
//	log.Info("An error occurred", sl.Err(err))
//
// In the example above, the error message "something went wrong" would be logged with the key "error".
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
