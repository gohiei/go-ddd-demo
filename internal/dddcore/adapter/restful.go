package dddcore

import (
	"cypt/internal/dddcore"

	"github.com/gin-gonic/gin"
)

// RestfulOutputError represents the structure of the error response to be rendered in a RESTful API.
type RestfulOutputError struct {
	Result     string `json:"result"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// RenderError renders the error response in a RESTful format.
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
