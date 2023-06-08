package user

import (
	"cypt/internal/dddcore"
	usecase "cypt/internal/user/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RenameUseCaseType dddcore.UseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput]
type RenameRestfulOutput struct {
	Result string                      `json:"result"`
	Ret    usecase.RenameUseCaseOutput `json:"ret"`
}

type RenameRestfulOutputError struct {
	Result     string `json:"result"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
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
		myerr := dddcore.NewErrorBy(err)

		ctx.Error(myerr)
		ctx.JSON(myerr.StatusCode, &RenameRestfulOutputError{
			Result:     "error",
			Message:    myerr.Message,
			Code:       myerr.Code,
			StatusCode: myerr.StatusCode,
		})

		return
	}

	ctx.JSON(http.StatusOK, &RenameRestfulOutput{
		Result: "ok",
		Ret:    output,
	})
}
