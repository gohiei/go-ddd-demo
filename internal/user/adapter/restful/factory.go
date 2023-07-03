package user

import (
	"github.com/gin-gonic/gin"

	"cypt/internal/dddcore"
	"cypt/internal/infra"
	adapter "cypt/internal/user/adapter"
	usecase "cypt/internal/user/usecase"
)

type UserRestfulConfig struct {
	UserWriteDatabaseDSN string
	UserReadDatabaseDSN  string
	IDRedisDSN           string
}

func NewUserRestful(router *gin.Engine, eventBus dddcore.EventBus, config UserRestfulConfig) {
	db, _ := infra.NewUserDB(
		config.UserWriteDatabaseDSN,
		config.UserReadDatabaseDSN,
	)
	userRepo := adapter.NewMySqlUserRepository(db)

	redisConn, _ := infra.NewIdRedis(config.IDRedisDSN)
	idRepo := adapter.NewRedisIDRepository(redisConn)

	usecase.NewNotifyManagerHandler(eventBus)
	NewRegisterUserRestful(router, usecase.NewRegisterUserUseCase(userRepo, idRepo, eventBus))
	NewRenameRestful(router, usecase.NewRenameUseCase(userRepo, eventBus))
}
