package user

import (
	dddcore "cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
	repo "cypt/internal/user/repository"
)

type RegisterUserUseCaseInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterUserUseCaseOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
}

type RegisterUserUseCase struct {
	userRepo repo.UserRepository
	idRepo   repo.IDRepository
	eventBus dddcore.EventBus
}

var _ dddcore.Input = (*RegisterUserUseCaseInput)(nil)
var _ dddcore.Output = (*RegisterUserUseCaseOutput)(nil)
var _ dddcore.UseCase[RegisterUserUseCaseInput, RegisterUserUseCaseOutput] = (*RegisterUserUseCase)(nil)

func NewRegisterUserUseCase(userRepo repo.UserRepository, idRepo repo.IDRepository, eb dddcore.EventBus) RegisterUserUseCase {
	return RegisterUserUseCase{
		userRepo: userRepo,
		idRepo:   idRepo,
		eventBus: eb,
	}
}

func (uc RegisterUserUseCase) Execute(input *RegisterUserUseCaseInput) (RegisterUserUseCaseOutput, error) {
	var user entity.User
	var userID int64
	var err error

	if userID, err = uc.idRepo.Incr(1); err != nil {
		return RegisterUserUseCaseOutput{}, err
	}

	if user, err = entity.NewUser(input.Username, input.Password, userID); err != nil {
		return RegisterUserUseCaseOutput{}, err
	}

	if err = uc.userRepo.Add(user); err != nil {
		return RegisterUserUseCaseOutput{}, err
	}

	uc.eventBus.PostAll(user)

	return RegisterUserUseCaseOutput{
		ID:       user.GetID().String(),
		Username: user.GetUsername(),
		UserID:   user.GetUserID(),
	}, nil
}
