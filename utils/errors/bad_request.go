package errors

import (
	"net/http"
	"runtime/debug"
)

// BadRequestError type
func BadRequestError(endpoint, method, message string) *Base {
	return NewError(
		"BadRequestError",
		endpoint,
		method,
		message,
		http.StatusBadRequest,
		debug.Stack(),
	)
}
