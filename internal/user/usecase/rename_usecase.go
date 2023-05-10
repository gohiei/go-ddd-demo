package user

import (
	dddcore "cypt/internal/dddcore"
	dto "cypt/internal/user/dto"
	repo "cypt/internal/user/repository"
)

type RenameUseCaseInput struct {
	Id       dddcore.UUID
	Username string
}

type RenameUseCaseOutput struct {
	Result string
	Ret    dto.UserDto
}

func (o *RenameUseCaseOutput) GetResult() string {
	return o.Result
}

func NewRenameUseCaseInput(id string, username string) RenameUseCaseInput {
	uuid, _ := dddcore.BuildUUID(id)

	return RenameUseCaseInput{Id: uuid, Username: username}
}

type RenameUseCase struct {
	userRepo repo.UserRepository
}

func NewRenameUseCase(repo repo.UserRepository) RenameUseCase {
	return RenameUseCase{
		userRepo: repo,
	}
}

func (useCase RenameUseCase) Execute(input *RenameUseCaseInput) (RenameUseCaseOutput, error) {
	user, err := useCase.userRepo.Get(input.Id)

	if err != nil {
		return RenameUseCaseOutput{}, err
	}

	user.Rename(input.Username)

	err = useCase.userRepo.Rename(user)

	if err != nil {
		return RenameUseCaseOutput{}, err
	}

	return RenameUseCaseOutput{
		Result: "ok",
		Ret:    dto.NewUserDto(user.GetId(), user.GetUsername()),
	}, nil
}
