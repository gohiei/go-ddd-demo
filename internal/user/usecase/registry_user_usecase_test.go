package user_test

import (
	"errors"
	"testing"

	usecase "cypt/internal/user/usecase"
	dddcoreMock "cypt/test/mocks/dddcore"
	userMock "cypt/test/mocks/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUserUseCase(t *testing.T) {
	r := userMock.NewUserRepository(t)
	r2 := userMock.NewIDRepository(t)
	b := dddcoreMock.NewEventBus(t)

	addFunc := r.On("Add", mock.Anything, mock.Anything).Return(nil)
	incrFunc := r2.On("Incr", 1).Return(int64(789), nil)
	postFunc := b.On("PostAll", mock.Anything).Return(nil)

	in := usecase.RegisterUserUseCaseInput{Username: "test1", Password: "password1"}
	uc := usecase.NewRegisterUserUseCase(r, r2, b)
	out, err := uc.Execute(&in)

	assert.Nil(t, err)
	assert.Equal(t, "test1", out.Username)
	assert.Equal(t, int64(789), out.UserID)
	assert.True(t, len(out.ID) > 0)

	r.AssertExpectations(t)
	r2.AssertExpectations(t)
	b.AssertExpectations(t)

	addFunc.Unset()
	incrFunc.Unset()
	postFunc.Unset()
}

func TestRegisterUserUseCaseWithErrFailedToAddUser(t *testing.T) {
	r := userMock.NewUserRepository(t)
	r2 := userMock.NewIDRepository(t)
	b := dddcoreMock.NewEventBus(t)

	addFunc := r.On("Add", mock.Anything, mock.Anything).Return(errors.New("fake error"))
	incrFunc := r2.On("Incr", 1).Return(int64(789), nil)

	in := usecase.RegisterUserUseCaseInput{Username: "test1", Password: "password1"}
	uc := usecase.NewRegisterUserUseCase(r, r2, b)
	_, err := uc.Execute(&in)

	assert.NotNil(t, err)

	r.AssertExpectations(t)
	r2.AssertExpectations(t)
	addFunc.Unset()
	incrFunc.Unset()
}
