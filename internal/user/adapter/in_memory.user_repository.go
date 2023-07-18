// Package adapter provides an implementation of the UserRepository interface
package adapter

import (
	"errors"
	"net/http"
	"sync"

	"cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
)

// InMemoryUserRepository is an implementation of `UserRepository` by memory
type InMemoryUserRepository struct {
	mutex sync.Mutex
	users map[dddcore.UUID]entity.User
}

// NewInMemoryUserRepository is a constructor of `InMemoryUserRepository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[dddcore.UUID]entity.User),
	}
}

// Get return `User` by a given id
func (repo *InMemoryUserRepository) Get(id dddcore.UUID) (entity.User, error) {
	if user, ok := repo.users[id]; ok {
		return user, nil
	}

	return entity.User{}, dddcore.NewErrorS("10006", "user not found", http.StatusBadRequest)
}

// Add `User` to collection
func (repo *InMemoryUserRepository) Add(u entity.User) error {
	if repo.users == nil {
		repo.mutex.Lock()
		repo.users = make(map[dddcore.UUID]entity.User)
		repo.mutex.Unlock()
	}

	if _, ok := repo.users[u.GetID()]; ok {
		return dddcore.NewErrorS("10008", "user already exists", http.StatusBadRequest)
	}

	repo.mutex.Lock()
	repo.users[u.GetID()] = u
	repo.mutex.Unlock()

	return nil
}

// Rename would replace the `User` entity
func (repo *InMemoryUserRepository) Rename(u entity.User) error {
	if _, ok := repo.users[u.GetID()]; !ok {
		return errors.New("user does not exist")
	}

	repo.mutex.Lock()
	repo.users[u.GetID()] = u
	repo.mutex.Unlock()

	return nil
}
