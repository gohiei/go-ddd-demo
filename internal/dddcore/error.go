package dddcore

import (
	"errors"
	"fmt"
	"net/http"
)

var _ error = (*Error)(nil)

// Error is a custom error type that implements the error interface
type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Detail     string `json:"detail"`
	Previous   error  `json:"previous"`
}

// Error returns the error message of the Error instance
func (e Error) Error() string {
	return fmt.Sprintf("%s (%s)", e.Message, e.Code)
}

// WithStatusCode sets the HTTP status code for the Error instance
func WithStatusCode(statusCode int) func(*Error) {
	return func(e *Error) {
		e.StatusCode = statusCode
	}
}

// WithDetail sets the detail message for the Error instance
func WithDetail(detail string) func(*Error) {
	return func(e *Error) {
		e.Detail = detail
	}
}

// WithPrevious sets the previous error for the Error instance
func WithPrevious(err error) func(*Error) {
	return func(e *Error) {
		e.Previous = e
	}
}

// NewError creates a new Error instance with the given code, message, and optional options
func NewError(code string, message string, options ...func(*Error)) Error {
	e := Error{Code: code, Message: message}

	for _, option := range options {
		option(&e)
	}

	return e
}

// NewErrorS creates a new Error instance with the given code, message, and HTTP status code
func NewErrorS(code string, message string, statusCode int) Error {
	return NewError(code, message, WithStatusCode(statusCode))
}

// NewErrorI creates a new Error instance with the given code and message, and HTTP status code set to Internal Server Error (500)
func NewErrorI(code string, message string) Error {
	return NewError(code, message, WithStatusCode(http.StatusInternalServerError))
}

// NewErrorBy creates a new Error instance based on an existing error, extracting information if it's already an Error or creating a new one with the error message
func NewErrorBy(err error) Error {
	var myerror Error

	if errors.As(err, &myerror) {
		return myerror
	}

	return NewErrorS("-", err.Error(), http.StatusInternalServerError)
}
