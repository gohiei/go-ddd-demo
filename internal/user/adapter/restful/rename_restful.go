package user

import (
	"net/http"

	"cypt/internal/dddcore"
	adapter "cypt/internal/dddcore/adapter"
	usecase "cypt/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

// RenameUseCaseType is a type alias for the RenameUseCase from dddcore package.
type RenameUseCaseType dddcore.UseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput]

// RenameRestfulOutput defines the output structure for the RenameRestful handler.
type RenameRestfulOutput struct {
	Result string                      `json:"result"`
	Ret    usecase.RenameUseCaseOutput `json:"ret"`
}

// NewRenameRestful registers the RenameRestful handler to the provided router with the given RenameUseCaseType.
func NewRenameRestful(router *gin.Engine, uc RenameUseCaseType) *RenameRestful {
	restful := &RenameRestful{Usecase: uc}
	router.PUT("/api/user/:id", restful.Execute)

	return restful
}

// RenameRestful is the handler for renaming a user.
type RenameRestful struct {
	Usecase RenameUseCaseType
}

// Execute is the handler function for renaming a user.
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
