package user

import (
	"cypt/internal/dddcore"
	"cypt/internal/infra"
	adapter "cypt/internal/user/adapter"
	usecase "cypt/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewUserRestful(router *gin.Engine, eventBus dddcore.EventBus, config *viper.Viper) {
	db, _ := infra.NewUserDB(
		config.GetString("user_write_db_dsn"),
		config.GetString("user_read_db_dsn"),
	)
	repo := adapter.NewMySqlUserRepository(db)

	NewRegisterUserRestful(router, usecase.NewRegisterUserUseCase(repo, eventBus))
	NewRenameRestful(router, usecase.NewRenameUseCase(repo, eventBus))
}
