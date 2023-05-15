package user

import (
	"cypt/internal/dddcore"
	repo "cypt/internal/user/repository"
	usecase "cypt/internal/user/usecase"
	"fmt"
)

type userController struct {
	registerUserUseCase  usecase.RegisterUserUseCase
	renameUseCase        usecase.RenameUseCase
	notifyManagerHandler usecase.NotifyManagerHandler
}

type UserController interface {
	RegisterUser(username string, password string) (usecase.RegisterUserUseCaseOutput, error)
	Rename(id string, username string) (usecase.RenameUseCaseOutput, error)
}

func NewController(r repo.UserRepository, eventBus dddcore.EventBus) UserController {
	return &userController{
		registerUserUseCase:  usecase.NewRegisterUserUseCase(r, eventBus),
		renameUseCase:        usecase.NewRenameUseCase(r, eventBus),
		notifyManagerHandler: usecase.NewNotifyManagerHandler(eventBus),
	}
}

func (c *userController) RegisterUser(username string, password string) (usecase.RegisterUserUseCaseOutput, error) {
	input := usecase.RegisterUserUseCaseInput{
		Username: username,
		Password: password,
	}

	out, err := c.registerUserUseCase.Execute(&input)

	if err != nil {
		return usecase.RegisterUserUseCaseOutput{}, err
	}

	fmt.Println(out.GetResult(), out.Ret, out.Ret.Username)

	return out, err
}

func (c *userController) Rename(id string, username string) (usecase.RenameUseCaseOutput, error) {
	input := usecase.NewRenameUseCaseInput(id, username)
	output, err := c.renameUseCase.Execute(&input)

	if err != nil {
		return usecase.RenameUseCaseOutput{}, err
	}

	fmt.Println(output.GetResult(), output.Ret.Username)

	return output, err
}
