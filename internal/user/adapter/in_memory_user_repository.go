package user

import (
	"fmt"
	"sync"

	"cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
	repository "cypt/internal/user/repository"
)

type InMemoryUserRepository struct {
	mutex sync.Mutex
	users map[dddcore.UUID]entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[dddcore.UUID]entity.User),
	}
}

func (repo *InMemoryUserRepository) Get(id dddcore.UUID) (entity.User, error) {
	if user, ok := repo.users[id]; ok {
		return user, nil
	}

	return entity.User{}, repository.ErrUserNotFound
}

func (repo *InMemoryUserRepository) Add(u entity.User) error {
	if repo.users == nil {
		repo.mutex.Lock()
		repo.users = make(map[dddcore.UUID]entity.User)
		repo.mutex.Unlock()
	}

	if _, ok := repo.users[u.GetId()]; ok {
		return fmt.Errorf("user already exists: %w", repository.ErrFailedToAddUser)
	}

	repo.mutex.Lock()
	repo.users[u.GetId()] = u
	repo.mutex.Unlock()

	return nil
}

func (repo *InMemoryUserRepository) Rename(u entity.User) error {
	if _, ok := repo.users[u.GetId()]; !ok {
		return fmt.Errorf("user does not exist: %w", repository.ErrFailedToRenameUser)
	}

	repo.mutex.Lock()
	repo.users[u.GetId()] = u
	repo.mutex.Unlock()

	return nil
}
