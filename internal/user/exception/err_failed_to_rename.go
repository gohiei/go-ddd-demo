package user

import "cypt/internal/dddcore"

type ErrFailedToRename struct {
	base *dddcore.Error
}

func NewErrFailedToRename() ErrFailedToRename {
	err := dddcore.NewError("10003", "failed to rename")
	return ErrFailedToRename{err}
}

func (e ErrFailedToRename) Error() string {
	return e.base.Error()
}

func (e ErrFailedToRename) With(err error) ErrFailedToRename {
	e.base.With(err)
	return e
}
