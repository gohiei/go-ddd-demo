package user_test

import (
	"cypt/internal/dddcore"
	user "cypt/internal/user/entity"
	repository "cypt/internal/user/repository"
	usecase "cypt/internal/user/usecase"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	dddcoreMock "cypt/test/mocks/dddcore"
	userMock "cypt/test/mocks/user"
)

func TestRenameUseCase(t *testing.T) {
	r := userMock.NewUserRepository(t)
	b := dddcoreMock.NewEventBus(t)

	uuid := dddcore.NewUUID()
	u := user.BuildUser(uuid.String(), "test2", "password2")

	getFunc := r.On("Get", mock.Anything).Return(u, nil)
	renameFunc := r.On("Rename", mock.Anything).Return(nil)
	postFunc := b.On("PostAll", mock.Anything).Return(nil)

	in := usecase.RenameUseCaseInput{Id: u.GetId(), Username: u.GetUsername()}
	uc := usecase.NewRenameUseCase(r, b)
	out, err := uc.Execute(&in)

	assert.Nil(t, err)
	assert.Equal(t, "ok", out.GetResult())
	assert.Equal(t, uuid.String(), out.Ret.Id)
	assert.Equal(t, "test2", out.Ret.Username)

	r.AssertExpectations(t)
	b.AssertExpectations(t)

	getFunc.Unset()
	renameFunc.Unset()
	postFunc.Unset()
}

func TestRenameUseCaseWithErrFailedToRenameUser(t *testing.T) {
	r := userMock.NewUserRepository(t)
	b := dddcoreMock.NewEventBus(t)

	uuid := dddcore.NewUUID()
	u := user.BuildUser(uuid.String(), "test2", "password2")

	getFunc := r.On("Get", mock.Anything).Return(u, nil)
	renameFunc := r.On("Rename", mock.Anything).Return(repository.ErrFailedToRenameUser)

	in := usecase.RenameUseCaseInput{Id: u.GetId(), Username: u.GetUsername()}
	uc := usecase.NewRenameUseCase(r, b)
	_, err := uc.Execute(&in)

	assert.NotNil(t, err)

	r.AssertExpectations(t)
	getFunc.Unset()
	renameFunc.Unset()
}
