package user

import (
	repo "cypt/internal/user/repository"
	usecase "cypt/internal/user/usecase"
	"fmt"
)

type userController struct {
	registryUserUseCase usecase.RegistryUserUseCase
	renameUseCase       usecase.RenameUseCase
}

type UserController interface {
	RegistryUser(username string, password string) (usecase.RegistryUserUseCaseOutput, error)
	Rename(id string, username string) (usecase.RenameUseCaseOutput, error)
}

func NewController(r repo.UserRepository) UserController {
	return &userController{
		registryUserUseCase: usecase.NewRegistryUserUseCase(r),
		renameUseCase:       usecase.NewRenameUseCase(r),
	}
}

func (c *userController) RegistryUser(username string, password string) (usecase.RegistryUserUseCaseOutput, error) {
	input := usecase.RegistryUserUseCaseInput{
		Username: username,
		Password: password,
	}

	out, err := c.registryUserUseCase.Execute(&input)

	if err != nil {
		return usecase.RegistryUserUseCaseOutput{}, err
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
