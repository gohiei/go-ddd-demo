package user

import "cypt/internal/dddcore"

type ErrUserNotFound struct {
	base *dddcore.Error
}

func NewErrUserNotFound() ErrUserNotFound {
	err := dddcore.NewError("10001", "user not found")
	return ErrUserNotFound{err}
}

func (e ErrUserNotFound) Error() string {
	return e.base.Error()
}

func (e ErrUserNotFound) With(err error) ErrUserNotFound {
	e.base.With(err)
	return e
}
