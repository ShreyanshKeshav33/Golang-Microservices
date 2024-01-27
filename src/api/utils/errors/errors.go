package errors

import "net/http"

type APIError interface {
	Status() int
	Message() string
	Error() string
}
type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AError   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}
func (e *apiError) Message() string {
	return e.AMessage
}
func (e *apiError) Error() string {
	return e.AError
}
func NewApiError(statusCode int, message string) APIError {
	return &apiError{
		AStatus:  statusCode,
		AMessage: message,
	}
}
func NewInternalServerError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}
func NewNotFoundError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}
func NewBadRequestError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
