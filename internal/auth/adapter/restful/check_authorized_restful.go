package auth

import (
	"fmt"

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

var (
	skips = map[string]bool{
		"POST /api/user": true,
	}
)

func (r *CheckAuthorizedRestful) Execute(ctx *gin.Context) {
	uri := fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL.Path)

	if skips[uri] {
		return
	}

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
