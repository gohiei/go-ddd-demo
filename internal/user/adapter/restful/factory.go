// Package restful provides RESTful API handlers for user-related operations.
package restful

import (
	"cypt/internal/dddcore"
	"cypt/internal/infra"
	logger "cypt/internal/logger/service"
	adapter "cypt/internal/user/adapter"
	usecase "cypt/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

// UserRestfulConfig holds the configuration for the user RESTful API.
type UserRestfulConfig struct {
	UserWriteDatabaseDSN string
	UserReadDatabaseDSN  string
	IDRedisDSN           string
}

// NewUserRestful sets up the user RESTful API routes and handlers.
func NewUserRestful(router *gin.Engine, eventBus dddcore.EventBus, config UserRestfulConfig) {
	db, _ := infra.NewUserDB(
		config.UserWriteDatabaseDSN,
		config.UserReadDatabaseDSN,
	)
	userRepo := adapter.NewMySQLUserRepository(db)

	redisConn, _ := infra.NewIDRedis(config.IDRedisDSN)
	idRepo := adapter.NewRedisIDRepository(redisConn)

	// Use a normal client
	// client := logger.NewHTTPClient(eventBus)

	// Use a client with decoder function
	client := logger.NewHTTPClientWithDecoder(eventBus, decoder)

	testAPIRepo := adapter.NewTestAPIOutsideRepository(client)

	usecase.NewNotifyManagerHandler(testAPIRepo, eventBus)
	NewRegisterUserRestful(router, usecase.NewRegisterUserUseCase(userRepo, idRepo, eventBus))
	NewRenameRestful(router, usecase.NewRenameUseCase(userRepo, eventBus))
}

// Example for decoding
func decoder(b []byte) (string, error) {
	body := "DECODE: " + string(b)
	return body, nil
}
