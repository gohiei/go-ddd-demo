package logger

import (
	"cypt/internal/dddcore"
	adapter "cypt/internal/logger/adapter"
	usecase "cypt/internal/logger/usecase"

	"github.com/gin-gonic/gin"
)

// NewLoggerRestful sets up the logger for a RESTful API using Gin.
// It configures the necessary Gin middlewares and sets up the log repository and use cases.
func NewLoggerRestful(router *gin.Engine, eventBus dddcore.EventBus, logDir string) {
	router.Use(RequestIDGenerator())
	router.Use(ErrorLogger())
	router.Use(NormalLogger())

	repo := adapter.NewZerologLogRepository(logDir)

	usecase.NewLogAccessUseCase(repo, eventBus)
	usecase.NewLogPostUseCase(repo, eventBus)
	usecase.NewLogErrorUseCase(repo, eventBus)
	usecase.NewLogHTTPRequestUseCase(repo, eventBus)
}
