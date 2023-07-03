package logger

import (
	"bytes"
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"cypt/internal/dddcore"
	events "cypt/internal/logger/entity/events"
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

// RequestIdGenerator is a Gin middleware that generates and adds a request ID to the request headers.
func RequestIdGenerator() gin.HandlerFunc {
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
			ctx.Request,
			&events.RequestDoneEventResponse{
				Latency:       latency,
				StatusCode:    cw.ResponseWriter.Status(),
				ContentLength: cw.ResponseWriter.Size(),
				ResponseData:  cw.buf.String(),
			},
		)

		if eb, _ := ctx.Get("event-bus"); eb != nil {
			eb.(dddcore.EventBus).Post(event)
		}
	}
}

// ErrorLogger is a Gin middleware that logs errors encountered during request processing.
func ErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		if eb, _ := ctx.Get("event-bus"); eb != nil {
			for _, curError := range ctx.Errors {
				var err dddcore.Error
				errors.As(curError.Err, &err)

				event := events.NewErrorRaisedEvent(
					start,
					ctx.ClientIP(),
					ctx.Request,
					err,
				)
				eb.(dddcore.EventBus).Post(event)
			}
		}
	}
}
