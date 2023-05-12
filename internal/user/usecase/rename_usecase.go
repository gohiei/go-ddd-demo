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
	eventBus dddcore.EventBus
}

func NewRenameUseCase(repo repo.UserRepository, eb dddcore.EventBus) RenameUseCase {
	return RenameUseCase{
		userRepo: repo,
		eventBus: eb,
	}
}

func (uc RenameUseCase) Execute(input *RenameUseCaseInput) (RenameUseCaseOutput, error) {
	user, err := uc.userRepo.Get(input.Id)

	if err != nil {
		return RenameUseCaseOutput{}, err
	}

	user.Rename(input.Username)

	err = uc.userRepo.Rename(user)

	if err != nil {
		return RenameUseCaseOutput{}, err
	}

	uc.eventBus.PostAll(user)

	return RenameUseCaseOutput{
		Result: "ok",
		Ret:    dto.NewUserDto(user.GetId(), user.GetUsername()),
	}, nil
}
