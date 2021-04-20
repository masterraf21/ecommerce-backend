package errors

import (
	"net/http"
	"runtime/debug"
)

// DefaultError type
func DefaultError(endpoint, method, message string) *Base {
	return NewError(
		"DefaultError",
		endpoint,
		method,
		message,
		http.StatusInternalServerError,
		debug.Stack(),
	)
}
