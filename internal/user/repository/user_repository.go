package user

import (
	"errors"

	"cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	Get(dddcore.UUID) (entity.User, error)
	Add(entity.User) error
	Rename(entity.User) error
}

var _ dddcore.Repository = (*UserRepository)(nil)
