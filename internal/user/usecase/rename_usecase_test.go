package user_test

import (
	"cypt/internal/dddcore"
	user "cypt/internal/user/entity"
	exception "cypt/internal/user/exception"
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

	in := usecase.RenameUseCaseInput{ID: u.GetID(), Username: u.GetUsername()}
	uc := usecase.NewRenameUseCase(r, b)
	out, err := uc.Execute(&in)

	assert.Nil(t, err)
	assert.Equal(t, uuid.String(), out.ID)
	assert.Equal(t, "test2", out.Username)

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
	renameFunc := r.On("Rename", mock.Anything).Return(exception.NewErrFailedToRename())

	in := usecase.RenameUseCaseInput{ID: u.GetID(), Username: u.GetUsername()}
	uc := usecase.NewRenameUseCase(r, b)
	_, err := uc.Execute(&in)

	assert.NotNil(t, err)
	assert.Error(t, exception.NewErrFailedToRename(), err)

	r.AssertExpectations(t)
	getFunc.Unset()
	renameFunc.Unset()
}
