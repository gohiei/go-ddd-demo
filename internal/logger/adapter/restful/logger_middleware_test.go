package logger_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"cypt/internal/dddcore"
	restful "cypt/internal/logger/adapter/restful"
	dddcoreMock "cypt/test/mocks/dddcore"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestErrorLoggerGivenExpectedError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/", new(bytes.Buffer))

	eb := dddcoreMock.NewEventBus(t)
	ctx.Set("event-bus", eb)
	ctx.Error(dddcore.NewErrorI("10xxx", "fake error 1"))

	handler := restful.ErrorLogger()
	handler(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"result\":\"error\",\"code\":\"10xxx\",\"message\":\"fake error 1\",\"request_id\":\"\",\"http_status_code\":500}", w.Body.String())
}

func TestErrorLoggerGivenUnexpectedError(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/", new(bytes.Buffer))

	eb := dddcoreMock.NewEventBus(t)
	ctx.Set("event-bus", eb)
	ctx.Error(errors.New("fake error 2"))

	postFunc := eb.On("Post", mock.Anything).Return(nil)

	handler := restful.ErrorLogger()
	handler(ctx)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"result\":\"error\",\"code\":\"-\",\"message\":\"fake error 2\",\"request_id\":\"\",\"http_status_code\":500}", w.Body.String())

	eb.AssertExpectations(t)
	postFunc.Unset()
}
