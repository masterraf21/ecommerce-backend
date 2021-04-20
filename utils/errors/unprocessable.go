package errors

import (
	"net/http"
	"runtime/debug"
)

// UnprocessableError type
func UnprocessableError(endpoint, method, message string) *Base {
	return NewError(
		"UnprocessableError",
		endpoint,
		method,
		message,
		http.StatusUnprocessableEntity,
		debug.Stack(),
	)
}
