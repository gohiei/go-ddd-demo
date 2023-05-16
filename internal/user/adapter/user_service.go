package user

import (
	"cypt/internal/dddcore"
	usecase "cypt/internal/user/usecase"
	"fmt"
)

type UserService struct {
	RegisterUserUseCase  dddcore.UseCase[usecase.RegisterUserUseCaseInput, usecase.RegisterUserUseCaseOutput]
	RenameUseCase        dddcore.UseCase[usecase.RenameUseCaseInput, usecase.RenameUseCaseOutput]
	NotifyManagerHandler usecase.NotifyManagerHandler
}

func (c *UserService) RegisterUser(username string, password string) (usecase.RegisterUserUseCaseOutput, error) {
	input := usecase.RegisterUserUseCaseInput{
		Username: username,
		Password: password,
	}

	out, err := c.RegisterUserUseCase.Execute(&input)

	if err != nil {
		return usecase.RegisterUserUseCaseOutput{}, err
	}

	fmt.Println(out.Result, out.Ret, out.Ret.Username)

	return out, err
}

func (c *UserService) Rename(id string, username string) (usecase.RenameUseCaseOutput, error) {
	input := usecase.NewRenameUseCaseInput(id, username)
	output, err := c.RenameUseCase.Execute(&input)

	if err != nil {
		return usecase.RenameUseCaseOutput{}, err
	}

	fmt.Println(output.Result, output.Ret.Username)

	return output, err
}
