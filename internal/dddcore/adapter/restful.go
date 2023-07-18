package dddcore

import (
	"github.com/gin-gonic/gin"
)

// RestfulOutputError represents the structure of the error response to be rendered in a RESTful API.
type RestfulOutputError struct {
	Result     string `json:"result"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	RequestID  string `json:"request_id"`
	StatusCode int    `json:"http_status_code"`
}

// RenderError renders the error response in a RESTful format.
func RenderError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	ctx.Abort()
}
