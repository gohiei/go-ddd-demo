package user

import (
	"cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
)

type UserRepository interface {
	Get(dddcore.UUID) (entity.User, error)
	Add(entity.User) error
	Rename(entity.User) error
}

var _ dddcore.Repository = (*UserRepository)(nil)
