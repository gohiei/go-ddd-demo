// Package user represents user bounded context
package user

import (
	"errors"
	"net/http"
	"time"

	"cypt/internal/dddcore"
	model "cypt/internal/user/adapter/model"
	entity "cypt/internal/user/entity"

	"gorm.io/gorm"
)

// MySQLUserRepository is an implementation of UserRepository using MySQL as the underlying storage.
type MySQLUserRepository struct {
	db *gorm.DB
}

// NewMySQLUserRepository creates a new instance of MySqlUserRepository.
func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db}
}

// Get retrieves a user entity by its ID from the MySQL database.
func (repo *MySQLUserRepository) Get(id dddcore.UUID) (entity.User, error) {
	user := model.UserModel{}
	result := repo.db.Take(&user, &model.UserModel{ID: id.String()})

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, dddcore.NewErrorS("10002", "user not found", http.StatusBadRequest)
		}

		return entity.User{}, dddcore.NewError(
			"10003", "failed to get",
			dddcore.WithStatusCode(http.StatusInternalServerError),
			dddcore.WithDetail(err.Error()),
			dddcore.WithPrevious(err),
		)
	}

	return entity.BuildUser(user.ID, user.Username, user.Password, user.UserID), nil
}

// Add adds a new user entity to the MySQL database.
func (repo *MySQLUserRepository) Add(u entity.User) error {
	user := model.UserModel{
		ID:       u.GetID().String(),
		Username: u.GetUsername(),
		Password: u.GetPassword(),
	}

	if result := repo.db.Create(&user); result.Error != nil {
		return dddcore.NewError(
			"10004", "failed to add",
			dddcore.WithStatusCode(http.StatusInternalServerError),
			dddcore.WithDetail(result.Error.Error()),
			dddcore.WithPrevious(result.Error),
		)
	}

	return nil
}

// Rename updates the username of an existing user entity in the MySQL database.
func (repo *MySQLUserRepository) Rename(u entity.User) error {
	user := model.UserModel{ID: u.GetID().String()}

	result := repo.db.Take(&user)

	if result.Error != nil {
		return dddcore.NewErrorS("10005", "user not found", http.StatusBadRequest)
	}

	result = repo.db.Model(&user).Updates(model.UserModel{
		Username:  u.GetUsername(),
		UpdatedAt: time.Now(),
	})

	if err := result.Error; err != nil {
		return dddcore.NewError(
			"10004", "failed to rename",
			dddcore.WithStatusCode(http.StatusInternalServerError),
			dddcore.WithDetail(result.Error.Error()),
			dddcore.WithPrevious(result.Error),
		)
	}

	return nil
}
