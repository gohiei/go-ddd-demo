package auth

import (
	"github.com/gin-gonic/gin"

	usecase "cypt/internal/auth/usecase"
	"cypt/internal/dddcore"
)

func NewAuthRestful(router *gin.Engine, eventBus dddcore.EventBus) {
	uc := usecase.NewCheckAuthorizationUsecase(eventBus)
	NewCheckAuthorizedRestful(router, uc)
}
