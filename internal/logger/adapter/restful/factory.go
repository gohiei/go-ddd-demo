package logger

import (
	"github.com/gin-gonic/gin"

	"cypt/internal/dddcore"
	adapter "cypt/internal/logger/adapter"
	usecase "cypt/internal/logger/usecase"
)

func NewLoggerRestful(router *gin.Engine, eventBus dddcore.EventBus, logDir string) {
	router.Use(RequestIdGenerator())
	router.Use(ErrorLogger())
	router.Use(NormalLogger())

	repo := adapter.NewZerologLogRepository(logDir)

	usecase.NewLogAccessUseCase(repo, eventBus)
	usecase.NewLogPostUseCase(repo, eventBus)
	usecase.NewLogErrorUseCase(repo, eventBus)
}
