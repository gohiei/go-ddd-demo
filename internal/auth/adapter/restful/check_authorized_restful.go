package auth

import (
	usecase "cypt/internal/auth/usecase"
	adapter "cypt/internal/dddcore/adapter"

	"github.com/gin-gonic/gin"
)

// NewCheckAuthorizedRestful registers the CheckAuthorizedRestful middleware to the provided router with the given CheckAuthorizationUsecase.
func NewCheckAuthorizedRestful(router *gin.Engine, usecase *usecase.CheckAuthorizationUsecase) *CheckAuthorizedRestful {
	restful := &CheckAuthorizedRestful{
		usecase: usecase,
	}

	router.Use(restful.Execute)

	return restful
}

// CheckAuthorizedRestful is a middleware for checking user authorization.
type CheckAuthorizedRestful struct {
	usecase *usecase.CheckAuthorizationUsecase
}

// Execute is the middleware function for checking user authorization.
func (r *CheckAuthorizedRestful) Execute(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	input := &usecase.CheckAuthorizationUsecaseInput{
		Token:  token,
		IP:     ctx.ClientIP(),
		XFF:    ctx.GetHeader("X-Forwarded-For"),
		Method: ctx.Request.Method,
		URL:    ctx.Request.URL.Path,
	}

	_, err := r.usecase.Execute(input)

	if err != nil {
		adapter.RenderError(ctx, err)
	}
}
