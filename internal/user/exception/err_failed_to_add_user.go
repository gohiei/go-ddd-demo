package user

import "cypt/internal/dddcore"

type ErrFailedToAddUser struct {
	base *dddcore.Error
}

func NewErrFailedToAddUser() ErrFailedToAddUser {
	err := dddcore.NewError("10002", "failed to add user")
	return ErrFailedToAddUser{err}
}

func (e ErrFailedToAddUser) Error() string {
	return e.base.Error()
}

func (e ErrFailedToAddUser) With(err error) ErrFailedToAddUser {
	e.base.With(err)
	return e
}
