package user_test

import (
	usecase "cypt/internal/user/usecase"

	repository "cypt/internal/user/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	dddcoreMock "cypt/test/mocks/dddcore"
	userMock "cypt/test/mocks/user"
)

func TestRegistryUserUseCase(t *testing.T) {
	r := userMock.NewUserRepository(t)
	b := dddcoreMock.NewEventBus(t)

	addFunc := r.On("Add", mock.Anything, mock.Anything).Return(nil)
	postFunc := b.On("PostAll", mock.Anything).Return(nil)

	in := usecase.RegistryUserUseCaseInput{Username: "test1", Password: "password1"}
	uc := usecase.NewRegistryUserUseCase(r, b)
	out, err := uc.Execute(&in)

	assert.Nil(t, err)
	assert.Equal(t, "ok", out.GetResult())
	assert.Equal(t, "test1", out.Ret.Username)
	assert.True(t, len(out.Ret.Id) > 0)

	r.AssertExpectations(t)
	b.AssertExpectations(t)

	addFunc.Unset()
	postFunc.Unset()
}

func TestRegistryUserUseCaseWithErrFailedToAddUser(t *testing.T) {
	r := userMock.NewUserRepository(t)
	b := dddcoreMock.NewEventBus(t)

	addFunc := r.On("Add", mock.Anything, mock.Anything).Return(repository.ErrFailedToAddUser)

	in := usecase.RegistryUserUseCaseInput{Username: "test1", Password: "password1"}
	uc := usecase.NewRegistryUserUseCase(r, b)
	_, err := uc.Execute(&in)

	assert.NotNil(t, err)

	r.AssertExpectations(t)
	addFunc.Unset()
}
