package errors

import (
	"net/http"
	"runtime/debug"
)

// NotFoundError type
func NotFoundError(endpoint, method, message string) *Base {
	return NewError(
		"NotFoundError",
		endpoint,
		method,
		message,
		http.StatusNotFound,
		debug.Stack(),
	)
}
