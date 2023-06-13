package user

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"cypt/internal/dddcore"
	model "cypt/internal/user/adapter/model"
	entity "cypt/internal/user/entity"
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
			return entity.User{}, dddcore.NewErrorS("10002", "user not found", http.StatusBadRequest)
		}

		return entity.User{}, dddcore.NewError(
			"10003", "failed to get",
			dddcore.WithStatusCode(http.StatusInternalServerError),
			dddcore.WithDetail(err.Error()),
			dddcore.WithPrevious(err),
		)
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
		return dddcore.NewError(
			"10004", "failed to add",
			dddcore.WithStatusCode(http.StatusInternalServerError),
			dddcore.WithDetail(result.Error.Error()),
			dddcore.WithPrevious(result.Error),
		)
	}

	return nil
}

func (repo *MySqlUserRepository) Rename(u entity.User) error {
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
