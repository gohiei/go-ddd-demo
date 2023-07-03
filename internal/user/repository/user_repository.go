// Package user represents user bounded context
package user

import (
	"cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
)

// UserRepository represents a repository for managing user entities.
type UserRepository interface {
	// Get retrieves a user entity by its ID.
	Get(dddcore.UUID) (entity.User, error)

	// Add adds a new user entity to the repository.
	Add(entity.User) error

	// Rename updates the username of an existing user entity.
	Rename(entity.User) error
}

var _ dddcore.Repository = (*UserRepository)(nil)
