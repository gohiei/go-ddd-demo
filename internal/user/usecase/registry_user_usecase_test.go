package user_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	usecase "cypt/internal/user/usecase"
	dddcoreMock "cypt/test/mocks/dddcore"
	userMock "cypt/test/mocks/user"
)

func TestRegisterUserUseCase(t *testing.T) {
	r := userMock.NewUserRepository(t)
	b := dddcoreMock.NewEventBus(t)

	addFunc := r.On("Add", mock.Anything, mock.Anything).Return(nil)
	postFunc := b.On("PostAll", mock.Anything).Return(nil)

	in := usecase.RegisterUserUseCaseInput{Username: "test1", Password: "password1"}
	uc := usecase.NewRegisterUserUseCase(r, b)
	out, err := uc.Execute(&in)

	assert.Nil(t, err)
	assert.Equal(t, "test1", out.Username)
	assert.True(t, len(out.ID) > 0)

	r.AssertExpectations(t)
	b.AssertExpectations(t)

	addFunc.Unset()
	postFunc.Unset()
}

func TestRegisterUserUseCaseWithErrFailedToAddUser(t *testing.T) {
	r := userMock.NewUserRepository(t)
	b := dddcoreMock.NewEventBus(t)

	addFunc := r.On("Add", mock.Anything, mock.Anything).Return(errors.New("fake error"))

	in := usecase.RegisterUserUseCaseInput{Username: "test1", Password: "password1"}
	uc := usecase.NewRegisterUserUseCase(r, b)
	_, err := uc.Execute(&in)

	assert.NotNil(t, err)

	r.AssertExpectations(t)
	addFunc.Unset()
}
