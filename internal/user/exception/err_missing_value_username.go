package user

import "cypt/internal/dddcore"

type ErrMissingValueUsername struct {
	base *dddcore.Error
}

func NewErrMissingValueUsername() ErrMissingValueUsername {
	err := dddcore.NewError("10004", "missing value `username`")
	return ErrMissingValueUsername{err}
}

func (e ErrMissingValueUsername) Error() string {
	return e.base.Error()
}

func (e ErrMissingValueUsername) With(err error) ErrMissingValueUsername {
	e.base.With(err)
	return e
}
