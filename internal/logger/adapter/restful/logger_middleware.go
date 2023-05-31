package logger

import (
	"bytes"
	"cypt/internal/dddcore"
	events "cypt/internal/logger/entity/events"
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

func Logger(eventBus dddcore.EventBus) gin.HandlerFunc {
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
