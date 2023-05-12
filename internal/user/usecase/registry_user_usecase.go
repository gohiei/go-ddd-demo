package user

import (
	dddcore "cypt/internal/dddcore"
	dto "cypt/internal/user/dto"
	entity "cypt/internal/user/entity"
	repo "cypt/internal/user/repository"
)

type RegistryUserUseCaseInput struct {
	Username string
	Password string
}

type RegistryUserUseCaseOutput struct {
	Result string
	Ret    dto.UserDto
}

func (out *RegistryUserUseCaseOutput) GetResult() string {
	return out.Result
}

type RegistryUserUseCase struct {
	userRepo repo.UserRepository
	eventBus dddcore.EventBus
}

var _ dddcore.Input = (*RegistryUserUseCaseInput)(nil)
var _ dddcore.Output = (*RegistryUserUseCaseOutput)(nil)

func NewRegistryUserUseCase(repo repo.UserRepository, eb dddcore.EventBus) RegistryUserUseCase {
	return RegistryUserUseCase{
		userRepo: repo,
		eventBus: eb,
	}
}

func (uc *RegistryUserUseCase) Execute(input *RegistryUserUseCaseInput) (RegistryUserUseCaseOutput, error) {
	user, err := entity.NewUser(input.Username, input.Password)

	if err != nil {
		return RegistryUserUseCaseOutput{}, err
	}

	err = uc.userRepo.Add(user)

	if err != nil {
		return RegistryUserUseCaseOutput{}, err
	}

	uc.eventBus.PostAll(user)

	return RegistryUserUseCaseOutput{
		Result: "ok",
		Ret:    dto.NewUserDto(user.GetId(), user.GetUsername()),
	}, nil
}
