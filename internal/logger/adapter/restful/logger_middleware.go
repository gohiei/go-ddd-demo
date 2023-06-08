package logger

import (
	"bytes"
	"cypt/internal/dddcore"
	events "cypt/internal/logger/entity/events"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

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

func NormalLogger(eventBus dddcore.EventBus) gin.HandlerFunc {
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

		eventBus.Post(event)
	}
}

func ErrorLogger(eb dddcore.EventBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		if len(ctx.Errors) > 0 {
			for _, curError := range ctx.Errors {
				var err dddcore.Error
				errors.As(curError.Err, &err)

				event := events.NewErrorRaisedEvent(
					start,
					ctx.ClientIP(),
					ctx.Request,
					err,
				)
				eb.Post(event)
			}
		}
	}
}
