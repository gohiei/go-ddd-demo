package dddcore

import (
	"fmt"
)

type Error struct {
	Code    string
	Message string
	Origin  error
}

func (e *Error) Error() string {
	if e.Origin != nil {
		return fmt.Sprintf("%s (%s): %s", e.Message, e.Code, e.Origin.Error())
	}

	return fmt.Sprintf("%s (%s)", e.Message, e.Code)
}

func (e *Error) With(err error) *Error {
	e.Origin = err
	return e
}

// func (e Error) Equal(err error) bool {
// 	if errors.Is(err, Error{}) {
// 		return e.Code == err.Code && e.Message == err.Message
// 	}
// }

func (e *Error) Equal(err error) bool {
	return e.Error() == err.Error()
}

func NewError(code string, message string) *Error {
	err := &Error{
		Code:    code,
		Message: message,
	}

	return err
}
