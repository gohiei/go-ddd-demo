package restful_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"cypt/internal/dddcore"
	adapter "cypt/internal/dddcore/adapter"
	logger "cypt/internal/logger/adapter/restful"
	restful "cypt/internal/user/adapter/restful"
	usecase "cypt/internal/user/usecase"
	dddcoreMock "cypt/test/mocks/dddcore"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getRouter(uc restful.RenameUseCaseType) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(logger.ErrorLogger())
	restful.NewRenameRestful(router, uc)

	return router
}

func TestRename(t *testing.T) {
	userID := dddcore.NewUUID().String()
	testcases := []struct {
		output *usecase.RenameUseCaseOutput
		err    error
		code   int
		result string
	}{
		{
			output: &usecase.RenameUseCaseOutput{
				ID:       userID,
				Username: "test2",
			},
			code:   http.StatusOK,
			result: "ok",
		},
		{
			output: nil,
			err:    dddcore.NewErrorS("10xxx", "user not found", http.StatusBadRequest),
			code:   http.StatusOK,
			result: "error",
		},
		{
			output: nil,
			err:    errors.New("other error"),
			code:   http.StatusInternalServerError,
			result: "error",
		},
	}

	for _, tc := range testcases {
		uc := dddcoreMock.NewUseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput](t)
		uc.On("Execute", mock.Anything).Return(tc.output, tc.err)

		router := getRouter(uc)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/user/"+userID, strings.NewReader("username=test2"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.code, w.Code)

		var out adapter.RestfulOutputError
		_ = json.Unmarshal(w.Body.Bytes(), &out)

		assert.Equal(t, tc.result, out.Result)
	}
}
