package user

import (
	"cypt/internal/dddcore"
	repository "cypt/internal/user/repository"
	usecase "cypt/internal/user/usecase"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RenameUseCaseType dddcore.UseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput]
type RenameRestfulOutput struct {
	Result string                      `json:"result"`
	Ret    usecase.RenameUseCaseOutput `json:"ret"`
}

type RenameRestfulOutputError struct {
	Result  string `json:"result"`
	Message string `json:"message"`
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
	var input usecase.RenameUseCaseInput
	ctx.Bind(&input)

	output, err := c.Usecase.Execute(&input)

	if err != nil {
		code := http.StatusInternalServerError

		if errors.Is(err, repository.ErrUserNotFound) {
			code = http.StatusBadRequest
		}

		ctx.JSON(code, &RenameRestfulOutputError{
			Result:  "error",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, &RenameRestfulOutput{
		Result: "ok",
		Ret:    output,
	})
}
