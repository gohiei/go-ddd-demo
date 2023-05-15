package user

import (
	dddcore "cypt/internal/dddcore"
	dto "cypt/internal/user/dto"
	entity "cypt/internal/user/entity"
	repo "cypt/internal/user/repository"
)

type RegisterUserUseCaseInput struct {
	Username string
	Password string
}

type RegisterUserUseCaseOutput struct {
	Result string
	Ret    dto.UserDto
}

type RegisterUserUseCase struct {
	userRepo repo.UserRepository
	eventBus dddcore.EventBus
}

var _ dddcore.Input = (*RegisterUserUseCaseInput)(nil)
var _ dddcore.Output = (*RegisterUserUseCaseOutput)(nil)

// var _ dddcore.UseCase[RegisterUserUseCaseInput, RegisterUserUseCaseOutput] = (*RegisterUserUseCase)(nil)

func NewRegisterUserUseCase(repo repo.UserRepository, eb dddcore.EventBus) RegisterUserUseCase {
	return RegisterUserUseCase{
		userRepo: repo,
		eventBus: eb,
	}
}

func (uc RegisterUserUseCase) Execute(input *RegisterUserUseCaseInput) (RegisterUserUseCaseOutput, error) {
	user, err := entity.NewUser(input.Username, input.Password)

	if err != nil {
		return RegisterUserUseCaseOutput{}, err
	}

	err = uc.userRepo.Add(user)

	if err != nil {
		return RegisterUserUseCaseOutput{}, err
	}

	uc.eventBus.PostAll(user)

	return RegisterUserUseCaseOutput{
		Result: "ok",
		Ret:    dto.NewUserDto(user.GetId(), user.GetUsername()),
	}, nil
}
