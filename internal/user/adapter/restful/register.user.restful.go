package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"cypt/internal/dddcore"
	usecase "cypt/internal/user/usecase"
)

func NewRegisterUserRestful(router *gin.Engine, uc usecase.RegisterUserUseCase) *RegisterUserRestful {
	restful := &RegisterUserRestful{Usecase: uc}
	router.POST("/api/user", restful.Execute)

	return restful
}

type RegisterUserRestful struct {
	Usecase dddcore.UseCase[usecase.RegisterUserUseCaseInput, usecase.RegisterUserUseCaseOutput]
}

func (c *RegisterUserRestful) Execute(ctx *gin.Context) {
	var input usecase.RegisterUserUseCaseInput
	ctx.Bind(&input)

	out, err := c.Usecase.Execute(&input)

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
