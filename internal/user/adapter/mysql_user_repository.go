package user

import (
	"cypt/internal/dddcore"
	model "cypt/internal/user/adapter/model"
	entity "cypt/internal/user/entity"
	repository "cypt/internal/user/repository"

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
	result := repo.db.Take(&user, &model.UserModel{Id: id.String()})

	if result.Error != nil {
		return entity.User{}, repository.ErrUserNotFound
	}

	return entity.BuildUser(user.Id, user.Username, user.Password), nil
}

func (repo *MySqlUserRepository) Add(u entity.User) error {
	user := model.UserModel{
		Id:       u.GetId().String(),
		Username: u.GetUsername(),
		Password: u.GetPassword(),
	}

	result := repo.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *MySqlUserRepository) Rename(u entity.User) error {
	user := model.UserModel{Id: u.GetId().String()}

	result := repo.db.Take(&user)

	if result.Error != nil {
		return repository.ErrUserNotFound
	}

	user.Username = u.GetUsername()

	repo.db.Save(&user)

	return nil
}
