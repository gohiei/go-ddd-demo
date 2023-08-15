package restful

import (
	"net/http"

	"cypt/internal/dddcore"
	"cypt/internal/dddcore/adapter"
	"cypt/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

// NewRegisterUserRestful creates and registers a new RESTful endpoint for user registration.
func NewRegisterUserRestful(router *gin.Engine, uc *usecase.RegisterUserUseCase) *RegisterUserRestful {
	restful := &RegisterUserRestful{Usecase: uc}
	router.POST("/api/user", restful.Execute)

	return restful
}

// RegisterUserRestful handles the user registration RESTful endpoint.
type RegisterUserRestful struct {
	Usecase dddcore.UseCase[usecase.RegisterUserUseCaseInput, usecase.RegisterUserUseCaseOutput]
}

type RegisterUserRestfulInput usecase.RegisterUserUseCaseInput

type RegisterUserRestfulOutput struct {
	Result string                             `json:"result"`
	Ret    *usecase.RegisterUserUseCaseOutput `json:"ret"`
}

// Execute handles the HTTP request for user registration.
// @Description Register User
// @Tags User
// @Accept json,x-www-form-urlencoded
// @Produce json
// @Param user body RegisterUserRestfulInput true "User Info"
// @Success 200 {object} RegisterUserRestfulOutput
// @Router /api/user [post]
func (c *RegisterUserRestful) Execute(ctx *gin.Context) {
	var input usecase.RegisterUserUseCaseInput

	if err := ctx.Bind(&input); err != nil {
		adapter.RenderError(ctx, err)
		return
	}

	out, err := c.Usecase.Execute(&input)

	if err != nil {
		adapter.RenderError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &RegisterUserRestfulOutput{
		Result: "ok",
		Ret:    out,
	})
}
