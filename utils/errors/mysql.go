package errors

import (
	"net/http"
	"runtime/debug"
)

// MySQLError type
func MySQLError(endpoint, method, message string) *Base {
	return NewError(
		"MySQLError",
		endpoint,
		method,
		message,
		http.StatusInternalServerError,
		debug.Stack(),
	)
}
