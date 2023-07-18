// Package repository provides interfaces for managing user entities
package repository

import (
	"cypt/internal/dddcore"
	"cypt/internal/user/entity"
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
