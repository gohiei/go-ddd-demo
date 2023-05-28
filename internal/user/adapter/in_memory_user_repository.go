package user

import (
	"errors"
	"sync"

	"cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
	exception "cypt/internal/user/exception"
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

	return entity.User{}, exception.NewErrUserNotFound()
}

func (repo *InMemoryUserRepository) Add(u entity.User) error {
	if repo.users == nil {
		repo.mutex.Lock()
		repo.users = make(map[dddcore.UUID]entity.User)
		repo.mutex.Unlock()
	}

	if _, ok := repo.users[u.GetID()]; ok {
		return exception.NewErrFailedToAddUser().With(errors.New("user already exists"))
	}

	repo.mutex.Lock()
	repo.users[u.GetID()] = u
	repo.mutex.Unlock()

	return nil
}

func (repo *InMemoryUserRepository) Rename(u entity.User) error {
	if _, ok := repo.users[u.GetID()]; !ok {
		return exception.NewErrFailedToRename().With(errors.New("user does not exist"))
	}

	repo.mutex.Lock()
	repo.users[u.GetID()] = u
	repo.mutex.Unlock()

	return nil
}
