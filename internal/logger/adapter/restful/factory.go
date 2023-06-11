package logger

import (
	"cypt/internal/dddcore"
	adapter "cypt/internal/logger/adapter"
	usecase "cypt/internal/logger/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewLoggerRestful(router *gin.Engine, eventBus dddcore.EventBus, config *viper.Viper) {
	router.Use(ErrorLogger())
	router.Use(NormalLogger())

	repo := adapter.NewZerologLogRepository(config.GetString("log_dir"))

	usecase.NewLogAccessUseCase(repo, eventBus)
	usecase.NewLogPostUseCase(repo, eventBus)
	usecase.NewLogErrorUseCase(repo, eventBus)
}
