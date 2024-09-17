package api

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {

	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func httpError(code int, message string) error {
	return &HTTPError{Code: code, Message: message}
}

func badRequest(message string) error {
	return httpError(http.StatusBadRequest, message)
}

func notFound(message string) error {
	return httpError(http.StatusNotFound, message)
}

func internalServerError(message string) error {
	return httpError(http.StatusInternalServerError, message)
}
