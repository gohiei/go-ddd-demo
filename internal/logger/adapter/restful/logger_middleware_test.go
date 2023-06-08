package logger_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"cypt/internal/dddcore"
	restful "cypt/internal/logger/adapter/restful"
	dddcoreMock "cypt/test/mocks/dddcore"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestNormalLogger(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest("GET", "/", new(bytes.Buffer))

	eb := dddcoreMock.NewEventBus(t)
	ctx.Set("event-bus", eb)

	postFunc := eb.On("Post", mock.Anything).Return(nil)

	handler := restful.NormalLogger()
	handler(ctx)

	eb.AssertExpectations(t)
	postFunc.Unset()
}

func TestErrorLogger(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest("GET", "/", new(bytes.Buffer))

	eb := dddcoreMock.NewEventBus(t)
	ctx.Set("event-bus", eb)
	ctx.Error(dddcore.NewErrorI("10xxx", "fake error 1"))
	ctx.Error(dddcore.NewErrorI("10xxy", "fake error 2"))

	postFunc := eb.On("Post", mock.Anything).Return(nil)

	handler := restful.ErrorLogger()
	handler(ctx)

	eb.AssertExpectations(t)
	postFunc.Unset()
}
