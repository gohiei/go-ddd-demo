package logger

import (
	"cypt/internal/dddcore"
	adapter "cypt/internal/logger/adapter"
	usecase "cypt/internal/logger/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoggerRestful(router *gin.Engine, eventBus dddcore.EventBus) {
	router.Use(ErrorLogger())
	router.Use(NormalLogger())
	repo := adapter.NewZerologLogRepository("/Users/chuck/Documents/Dev/pineapple/pineapple-go-micro/logs")

	usecase.NewLogAccessUseCase(repo, eventBus)
	usecase.NewLogPostUseCase(repo, eventBus)
	usecase.NewLogErrorUseCase(repo, eventBus)
}
