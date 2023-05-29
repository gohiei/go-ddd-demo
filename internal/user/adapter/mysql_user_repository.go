package user

import (
	"cypt/internal/dddcore"
	model "cypt/internal/user/adapter/model"
	entity "cypt/internal/user/entity"
	repository "cypt/internal/user/repository"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MySqlUserRepository struct {
	db *gorm.DB
}

func NewMySqlUserRepository(db *gorm.DB) *MySqlUserRepository {
	return &MySqlUserRepository{db}
}

func (repo *MySqlUserRepository) Get(id dddcore.UUID) (entity.User, error) {
	user := model.UserModel{}
	result := repo.db.Take(&user, &model.UserModel{ID: id.String()})

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, repository.ErrUserNotFound
		}

		return entity.User{}, fmt.Errorf("failed to get by id `%s`: %w", id, err)
	}

	return entity.BuildUser(user.ID, user.Username, user.Password), nil
}

func (repo *MySqlUserRepository) Add(u entity.User) error {
	user := model.UserModel{
		ID:       u.GetID().String(),
		Username: u.GetUsername(),
		Password: u.GetPassword(),
	}

	if result := repo.db.Create(&user); result.Error != nil {
		return fmt.Errorf("failed to add: %w", result.Error)
	}

	return nil
}

func (repo *MySqlUserRepository) Rename(u entity.User) error {
	user := model.UserModel{ID: u.GetID().String()}

	result := repo.db.Take(&user)

	if result.Error != nil {
		return repository.ErrUserNotFound
	}

	result = repo.db.Model(&user).Updates(model.UserModel{
		Username:  u.GetUsername(),
		UpdatedAt: time.Now(),
	})

	if err := result.Error; err != nil {
		return fmt.Errorf("failed to rename `%s`: %w", user.ID, err)
	}

	return nil
}
