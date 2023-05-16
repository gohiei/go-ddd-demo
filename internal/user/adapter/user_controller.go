package user

import (
	"cypt/internal/dddcore"
	repo "cypt/internal/user/repository"
	usecase "cypt/internal/user/usecase"
)

type UserController interface {
	RegisterUser(username string, password string) (usecase.RegisterUserUseCaseOutput, error)
	Rename(id string, username string) (usecase.RenameUseCaseOutput, error)
}

func NewController(r repo.UserRepository, eventBus dddcore.EventBus) UserController {
	return &UserService{
		RegisterUserUseCase:  usecase.NewRegisterUserUseCase(r, eventBus),
		RenameUseCase:        usecase.NewRenameUseCase(r, eventBus),
		NotifyManagerHandler: usecase.NewNotifyManagerHandler(eventBus),
	}
}
