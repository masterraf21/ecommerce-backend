package http

import (
	"encoding/json"
	"net/http"

	"github.com/masterraf21/ecommerce-backend/utils/errors"
	logger "github.com/sirupsen/logrus"
)

// HandleError HTTP Response
func HandleError(w http.ResponseWriter, r *http.Request, err error, message string, statusCode int) {
	body := struct {
		Message string `json:"message"`
		Detail  string `json:"detail"`
	}{
		Message: message,
		Detail:  err.Error(),
	}

	switch statusCode {
	case http.StatusBadRequest:
		logger.Error(errors.BadRequestError(r.URL.EscapedPath(), r.Method, message+", "+err.Error()))
	case http.StatusUnprocessableEntity:
		logger.Error(errors.UnprocessableError(r.URL.EscapedPath(), r.Method, message+", "+err.Error()))
	case http.StatusNotFound:
		logger.Error(errors.NotFoundError(r.URL.EscapedPath(), r.Method, message+", "+err.Error()))
	case http.StatusInternalServerError:
		logger.Error(errors.DefaultError(r.URL.EscapedPath(), r.Method, message+", "+err.Error()))
	default:
		logger.Error(errors.DefaultError(r.URL.EscapedPath(), r.Method, message+", "+err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// HandleJSONResponse HTTP
func HandleJSONResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	message, err := json.Marshal(v)
	if err != nil {
		HandleError(w, r, err, "failed to marshal response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// HandleNoJSONResponse HTTP
func HandleNoJSONResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
