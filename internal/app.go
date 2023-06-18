package app

import (
	adapter "cypt/internal/dddcore/adapter"
	logger "cypt/internal/logger/adapter/restful"
	user "cypt/internal/user/adapter/restful"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewAppController(router *gin.Engine, config *viper.Viper) {
	eventBus := adapter.NewWatermillEventBus()

	router.Use(func(c *gin.Context) {
		c.Set("event-bus", &eventBus)
	})

	logger.NewLoggerRestful(router, &eventBus, config.GetString("log_dir"))
	user.NewUserRestful(router, &eventBus, user.UserRestfulConfig{
		UserWriteDatabaseDSN: config.GetString("user_write_db_dsn"),
		UserReadDatabaseDSN:  config.GetString("user_read_db_dsn"),
		IDRedisDSN:           config.GetString("id_redis_dsn"),
	})
}
