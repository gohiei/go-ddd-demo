package dddcore

import (
	"errors"
	"fmt"
	"net/http"
)

var _ error = (*Error)(nil)

type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Detail     string `json:"string"`
	Previous   error  `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s (%s)", e.Message, e.Code)
}

func WithStatusCode(statusCode int) func(*Error) {
	return func(e *Error) {
		e.StatusCode = statusCode
	}
}

func WithDetail(detail string) func(*Error) {
	return func(e *Error) {
		e.Detail = detail
	}
}

func WithPrevious(err error) func(*Error) {
	return func(e *Error) {
		e.Previous = e
	}
}

func NewError(code string, message string, options ...func(*Error)) Error {
	e := Error{Code: code, Message: message}

	for _, option := range options {
		option(&e)
	}

	return e
}

func NewErrorS(code string, message string, statusCode int) Error {
	return NewError(code, message, WithStatusCode(statusCode))
}

func NewErrorI(code string, message string) Error {
	return NewError(code, message, WithStatusCode(http.StatusInternalServerError))
}

func FormatBy(err error) (string, string, int) {
	var myerror Error

	if errors.As(err, &myerror) {
		return myerror.Code, myerror.Message, myerror.StatusCode
	}

	return "-", err.Error(), http.StatusInternalServerError
}
