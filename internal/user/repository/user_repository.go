package user

import (
	"errors"

	"cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
)

var (
	ErrUserNotFound       = errors.New("the user was not found in the repository")
	ErrFailedToAddUser    = errors.New("failed to add the user to the repository")
	ErrFailedToRenameUser = errors.New("failed to rename the user in the repository")
)

type UserRepository interface {
	Get(dddcore.UUID) (entity.User, error)
	Add(entity.User) error
	Rename(entity.User) error
}

var _ dddcore.Repository = (*UserRepository)(nil)
