package user

import (
	"cypt/internal/dddcore"
	"cypt/internal/infra"
	adapter "cypt/internal/user/adapter"
	usecase "cypt/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

func NewUserRestful(router *gin.Engine, eventBus dddcore.EventBus) {
	db, _ := infra.NewUserDB()
	repo := adapter.NewMySqlUserRepository(db)

	NewRegisterUserRestful(router, usecase.NewRegisterUserUseCase(repo, eventBus))
	NewRenameRestful(router, usecase.NewRenameUseCase(repo, eventBus))
}
