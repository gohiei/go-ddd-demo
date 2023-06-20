package auth

import (
	"github.com/gin-gonic/gin"

	usecase "cypt/internal/auth/usecase"
	adapter "cypt/internal/dddcore/adapter"
)

func NewCheckAuthorizedRestful(router *gin.Engine, usecase *usecase.CheckAuthorizationUsecase) *CheckAuthorizedRestful {
	restful := &CheckAuthorizedRestful{
		usecase: usecase,
	}

	router.Use(restful.Execute)

	return restful
}

type CheckAuthorizedRestful struct {
	usecase *usecase.CheckAuthorizationUsecase
}

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
