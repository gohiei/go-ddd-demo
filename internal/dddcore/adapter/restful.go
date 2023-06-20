package dddcore

import (
	"cypt/internal/dddcore"

	"github.com/gin-gonic/gin"
)

type RestfulOutputError struct {
	Result     string `json:"result"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func RenderError(ctx *gin.Context, err error) {
	myerr := dddcore.NewErrorBy(err)

	ctx.Error(myerr)
	ctx.JSON(myerr.StatusCode, &RestfulOutputError{
		Result:     "error",
		Message:    myerr.Message,
		Code:       myerr.Code,
		StatusCode: myerr.StatusCode,
	})
	ctx.Abort()
}
