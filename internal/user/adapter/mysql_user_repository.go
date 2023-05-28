package user

import (
	"cypt/internal/dddcore"
	model "cypt/internal/user/adapter/model"
	entity "cypt/internal/user/entity"
	exception "cypt/internal/user/exception"
	"errors"
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
			return entity.User{}, exception.NewErrUserNotFound()
		}

		return entity.User{}, exception.NewErrUserNotFound().With(err)
	}

	return entity.BuildUser(user.ID, user.Username, user.Password), nil
}

func (repo *MySqlUserRepository) Add(u entity.User) error {
	user := model.UserModel{
		ID:       u.GetID().String(),
		Username: u.GetUsername(),
		Password: u.GetPassword(),
	}

	result := repo.db.Create(&user)

	if err := result.Error; err != nil {
		return exception.NewErrFailedToAddUser().With(result.Error)
	}

	return nil
}

func (repo *MySqlUserRepository) Rename(u entity.User) error {
	user := model.UserModel{ID: u.GetID().String()}

	_, err := repo.Get(u.GetID())

	if err != nil {
		return err
	}

	result := repo.db.Model(&user).Updates(model.UserModel{
		Username:  u.GetUsername(),
		UpdatedAt: time.Now(),
	})

	if err := result.Error; err != nil {
		return exception.NewErrFailedToRename().With(result.Error)
	}

	return nil
}
