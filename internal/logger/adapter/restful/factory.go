// Package restful provides functions for setting up a logger for a RESTful API using Gin.
// It configures the necessary Gin middlewares and sets up the log repository and use cases.
package restful

import (
	"cypt/internal/dddcore"
	"cypt/internal/logger/adapter"
	"cypt/internal/logger/usecase"

	"github.com/gin-gonic/gin"
)

type LogConfig struct {
	Dir           string
	TrustedProxy  []string
	TrustedHeader string
}

// NewLoggerRestful sets up the logger for a RESTful API using Gin.
// It configures the necessary Gin middlewares and sets up the log repository and use cases.
func NewLoggerRestful(router *gin.Engine, eventBus dddcore.EventBus, config LogConfig) {
	router.Use(GetClientIP(router, config.TrustedProxy, config.TrustedHeader))
	router.Use(RequestIDGenerator())
	router.Use(NormalLogger())
	router.Use(ErrorLogger())

	// repo := adapter.NewZerologLogRepository(logDir)
	repo := adapter.NewZapLogRepository(config.Dir)

	usecase.NewLogAccessUseCase(repo, eventBus)
	usecase.NewLogPostUseCase(repo, eventBus)
	usecase.NewLogErrorUseCase(repo, eventBus)
	usecase.NewLogHTTPRequestUseCase(repo, eventBus)
}
