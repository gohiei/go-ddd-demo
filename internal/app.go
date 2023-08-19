// Package app builds a full application.
package app

import (
	auth "cypt/internal/auth/adapter/restful"
	dddcore "cypt/internal/dddcore/adapter"
	logger "cypt/internal/logger/adapter/restful"
	swagger "cypt/internal/swagger/adapter/restful"
	userGrpc "cypt/internal/user/adapter/grpc"
	userRestful "cypt/internal/user/adapter/restful"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// NewAppRestfulServer initializes the application controller.
func NewAppRestfulServer(router *gin.Engine, config *viper.Viper) {
	eventBus := dddcore.NewWatermillEventBus()

	router.Use(func(c *gin.Context) {
		c.Set("event-bus", &eventBus)
	})

	logger.NewLoggerRestful(router, &eventBus, logger.LogConfig{
		Dir:           config.GetString("log_dir"),
		TrustedHeader: config.GetString("trusted_header"),
		TrustedProxy:  config.GetStringSlice("trusted_proxy"),
	})

	auth.NewAuthRestful(router, &eventBus)

	userRestful.NewUserRestful(router, &eventBus, userRestful.UserRestfulConfig{
		UserWriteDatabaseDSN: config.GetString("user_write_db_dsn"),
		UserReadDatabaseDSN:  config.GetString("user_read_db_dsn"),
		IDRedisDSN:           config.GetString("id_redis_dsn"),
	})

	swagger.NewSwaggerRestful(router)
}

func NewAppGrpcServer(server *grpc.Server, config *viper.Viper) {
	eventBus := dddcore.NewWatermillEventBus()

	userGrpc.NewUserGrpc(server, &eventBus, userGrpc.UserGrpcConfig{
		UserWriteDatabaseDSN: config.GetString("user_write_db_dsn"),
		UserReadDatabaseDSN:  config.GetString("user_read_db_dsn"),
		IDRedisDSN:           config.GetString("id_redis_dsn"),
	})
}
