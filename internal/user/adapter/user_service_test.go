package user_test

import (
	"cypt/internal/dddcore"
	user "cypt/internal/user/adapter"
	usecase "cypt/internal/user/usecase"
	mocks "cypt/test/mocks/dddcore"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceRegisterUserUseCase(t *testing.T) {
	uc := mocks.NewUseCase[usecase.RegisterUserUseCaseInput, usecase.RegisterUserUseCaseOutput](t)

	executeFunc := uc.On("Execute", mock.Anything).Return(
		usecase.RegisterUserUseCaseOutput{Result: "ok"},
		nil,
	)

	s := user.UserService{RegisterUserUseCase: uc}
	output, err := s.RegisterUser("test1", "password1")

	assert.Nil(t, err)
	assert.NotEmpty(t, output.Result)

	uc.AssertExpectations(t)
	executeFunc.Unset()
}

func TestUserServiceRegisterUserUseCaseWithError(t *testing.T) {
	uc := mocks.NewUseCase[usecase.RegisterUserUseCaseInput, usecase.RegisterUserUseCaseOutput](t)

	executeFunc := uc.On("Execute", mock.Anything).Return(
		usecase.RegisterUserUseCaseOutput{},
		errors.New("fake test"),
	)

	s := user.UserService{RegisterUserUseCase: uc}
	output, err := s.RegisterUser("test1", "password1")

	assert.NotNil(t, err)
	assert.Empty(t, output.Result)

	uc.AssertExpectations(t)
	executeFunc.Unset()
}

func TestUserServiceRenameUseCase(t *testing.T) {
	uc := mocks.NewUseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput](t)

	executeFunc := uc.On("Execute", mock.Anything).Return(
		usecase.RenameUseCaseOutput{Result: "ok"},
		nil,
	)

	s := user.UserService{RenameUseCase: uc}
	output, err := s.Rename(dddcore.NewUUID().String(), "test2")

	assert.Nil(t, err)
	assert.NotEmpty(t, output.Result)

	uc.AssertExpectations(t)
	executeFunc.Unset()
}

func TestUserServiceRenameUseCaseWithError(t *testing.T) {
	uc := mocks.NewUseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput](t)

	executeFunc := uc.On("Execute", mock.Anything).Return(
		usecase.RenameUseCaseOutput{},
		errors.New("fake test"),
	)

	s := user.UserService{RenameUseCase: uc}
	output, err := s.Rename(dddcore.NewUUID().String(), "test2")

	assert.NotNil(t, err)
	assert.Empty(t, output.Result)

	uc.AssertExpectations(t)
	executeFunc.Unset()
}
