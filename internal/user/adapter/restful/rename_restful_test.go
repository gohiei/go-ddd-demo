package user_test

import (
	"cypt/internal/dddcore"
	restful "cypt/internal/user/adapter/restful"
	repository "cypt/internal/user/repository"
	usecase "cypt/internal/user/usecase"
	dddcoreMock "cypt/test/mocks/dddcore"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getRouter(uc restful.RenameUseCaseType) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	restful.NewRenameRestful(router, uc)

	return router
}

func TestRename(t *testing.T) {
	userId := dddcore.NewUUID().String()
	testcases := []struct {
		output usecase.RenameUseCaseOutput
		err    error
		code   int
		result string
	}{
		{
			output: usecase.RenameUseCaseOutput{
				ID:       userId,
				Username: "test2",
			},
			code:   http.StatusOK,
			result: "ok",
		},
		{
			output: usecase.RenameUseCaseOutput{},
			err:    repository.ErrUserNotFound,
			code:   http.StatusBadRequest,
			result: "error",
		},
		{
			output: usecase.RenameUseCaseOutput{},
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
		req, _ := http.NewRequest("PUT", "/api/user/"+userId, strings.NewReader("username=test2"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, tc.code, w.Code)

		var out restful.RenameRestfulOutputError
		err := json.Unmarshal(w.Body.Bytes(), &out)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(w.Body.String())
		assert.Equal(t, tc.result, out.Result)
	}
}
