package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"cypt/internal/dddcore"
	adapter "cypt/internal/dddcore/adapter"
	usecase "cypt/internal/user/usecase"
)

type RenameUseCaseType dddcore.UseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput]
type RenameRestfulOutput struct {
	Result string                      `json:"result"`
	Ret    usecase.RenameUseCaseOutput `json:"ret"`
}

func NewRenameRestful(router *gin.Engine, uc RenameUseCaseType) *RenameRestful {
	restful := &RenameRestful{Usecase: uc}
	router.PUT("/api/user/:id", restful.Execute)

	return restful
}

type RenameRestful struct {
	Usecase RenameUseCaseType
}

func (c *RenameRestful) Execute(ctx *gin.Context) {
	input := usecase.RenameUseCaseInput{
		ID:       ctx.Param("id"),
		Username: ctx.PostForm("username"),
	}

	output, err := c.Usecase.Execute(&input)

	if err != nil {
		adapter.RenderError(ctx, err)

		return
	}

	ctx.JSON(http.StatusOK, &RenameRestfulOutput{
		Result: "ok",
		Ret:    output,
	})
}
