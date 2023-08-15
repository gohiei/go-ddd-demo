// Package restful provides RESTful API handlers for user-related operations.
package restful

import (
	"net/http"

	"cypt/internal/dddcore"
	"cypt/internal/dddcore/adapter"
	"cypt/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

// RenameUseCaseType is a type alias for the RenameUseCase from dddcore package.
type RenameUseCaseType dddcore.UseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput]

// RenameRestfulOutput defines the output structure for the RenameRestful handler.
type RenameRestfulOutput struct {
	Result string                       `json:"result"`
	Ret    *usecase.RenameUseCaseOutput `json:"ret"`
}

type RenameRestfulInput struct {
	Username string `json:"usernmae"`
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
// @Description Rename a user
// @Tags User
// @Accept json,x-www-form-urlencoded
// @Produce json
// @Param id path string true "User ID"
// @Param user body RenameRestfulInput true "User"
// @Success 200 {object} RenameRestfulOutput
// @Router /api/user/:id [put]
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
