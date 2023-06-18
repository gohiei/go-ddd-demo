package user

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"cypt/internal/dddcore"
	"cypt/internal/infra"
	adapter "cypt/internal/user/adapter"
	usecase "cypt/internal/user/usecase"
)

func NewUserRestful(router *gin.Engine, eventBus dddcore.EventBus, config *viper.Viper) {
	db, _ := infra.NewUserDB(
		config.GetString("user_write_db_dsn"),
		config.GetString("user_read_db_dsn"),
	)
	userRepo := adapter.NewMySqlUserRepository(db)

	redisConn, _ := infra.NewIdRedis(config.GetString("id_redis_dsn"))
	idRepo := adapter.NewRedisIDRepository(redisConn)

	NewRegisterUserRestful(router, usecase.NewRegisterUserUseCase(userRepo, idRepo, eventBus))
	NewRenameRestful(router, usecase.NewRenameUseCase(userRepo, eventBus))
}
