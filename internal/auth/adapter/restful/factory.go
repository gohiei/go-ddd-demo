// Package restful provides functionality for setting up authentication-related RESTful endpoints.
package restful

import (
	"cypt/internal/auth/usecase"
	"cypt/internal/dddcore"

	"github.com/gin-gonic/gin"
)

// NewAuthRestful sets up the authentication-related RESTful endpoints.
func NewAuthRestful(router *gin.Engine, eventBus dddcore.EventBus) {
	uc := usecase.NewCheckAuthorizationUsecase(eventBus)
	NewCheckAuthorizedRestful(router, uc)
}
