package restful

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"cypt/internal/dddcore"
	"cypt/internal/dddcore/adapter"
	"cypt/internal/logger/entity/events"

	"github.com/gin-gonic/gin"
)

// copyWriter is a custom ResponseWriter that captures the response data.
type copyWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func (cw *copyWriter) Write(b []byte) (int, error) {
	if count, err := cw.buf.Write(b); err != nil {
		return count, err
	}

	return cw.ResponseWriter.Write(b)
}

// RequestIDGenerator is a Gin middleware that generates and adds a request ID to the request headers.
func RequestIDGenerator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if rid := ctx.GetHeader("X-Request-Id"); rid == "" {
			rid = dddcore.NewUUID().String()

			ctx.Request.Header.Set("X-Request-Id", rid)
			ctx.Header("X-Request-Id", rid)
		}

		ctx.Next()
	}
}

// NormalLogger is a Gin middleware that logs normal HTTP requests.
func NormalLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		cw := &copyWriter{
			ResponseWriter: ctx.Writer,
			buf:            &bytes.Buffer{},
		}
		ctx.Writer = cw

		ctx.Next()

		latency := time.Since(start)

		event := events.NewRequestDoneEvent(
			start,
			ctx.ClientIP(),
			ctx.FullPath(),
			ctx.Request,
			&events.RequestDoneEventResponse{
				Latency:       latency,
				StatusCode:    cw.ResponseWriter.Status(),
				ContentLength: cw.ResponseWriter.Size(),
				ResponseData:  cw.buf.String(),
			},
		)

		if eb, _ := ctx.Get("event-bus"); eb != nil {
			if !strings.HasPrefix(ctx.Request.RequestURI, "/api") {
				return
			}

			eb.(dddcore.EventBus).Post(event)
		}
	}
}

// ErrorLogger is a middleware function that handles error logging and response generation for Gin framework.
// It captures any errors occurred during the request processing and logs them using the defined format.
// If the error is an expected domain-specific error, it returns an appropriate JSON response with the corresponding status code.
// Otherwise, it creates a new error based on the original error and returns an internal server error response.
// Additionally, if an event bus is available in the context, it publishes an unexpected-error-raised event for further processing.
func ErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors[0].Err

		var cerr dddcore.Error
		var statusCode int

		_, isExpected := err.(dddcore.Error)

		if isExpected {
			cerr = err.(dddcore.Error)
			statusCode = http.StatusOK
		} else {
			cerr = dddcore.NewErrorBy(err)
			statusCode = http.StatusInternalServerError
		}

		ctx.JSON(statusCode, &adapter.RestfulOutputError{
			Result:     "error",
			Message:    cerr.Message,
			Code:       cerr.Code,
			RequestID:  ctx.Request.Header.Get("X-Request-ID"),
			StatusCode: cerr.StatusCode,
		})

		if isExpected {
			return
		}

		if eb, _ := ctx.Get("event-bus"); eb != nil {
			event := events.NewUnexpectedErrorRaisedEvent(
				start,
				ctx.ClientIP(),
				ctx.Request,
				cerr,
			)
			eb.(dddcore.EventBus).Post(event)
		}
	}
}
