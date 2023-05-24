package user

import (
	"cypt/internal/dddcore"
	usecase "cypt/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	RegisterUserUseCase  dddcore.UseCase[usecase.RegisterUserUseCaseInput, usecase.RegisterUserUseCaseOutput]
	RenameUseCase        dddcore.UseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput]
	NotifyManagerHandler usecase.NotifyManagerHandler
}

func (c *UserService) RegisterUser(ctx *gin.Context) {
	var input usecase.RegisterUserUseCaseInput
	ctx.Bind(&input)

	out, err := c.RegisterUserUseCase.Execute(&input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result":  "error",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "ok",
		"ret":    out,
	})
}

func (c *UserService) Rename(ctx *gin.Context) {
	var input usecase.RenameUseCaseInput
	ctx.Bind(&input)

	output, err := c.RenameUseCase.Execute(&input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result":  "error",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "ok",
		"ret":    output,
	})
}
