// Package usecase provides the business logic and use cases for managing user
package usecase

import (
	"cypt/internal/dddcore"
	"cypt/internal/user/entity"
	"cypt/internal/user/repository"
)

// RegisterUserUseCaseInput represents the input data for the RegisterUserUseCase.
type RegisterUserUseCaseInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// RegisterUserUseCaseOutput represents the output data for the RegisterUserUseCase.
type RegisterUserUseCaseOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	UserID   int64  `json:"user_id"`
}

// RegisterUserUseCase is a use case for registering a new user.
type RegisterUserUseCase struct {
	userRepo repository.UserRepository
	idRepo   repository.IDRepository
	eventBus dddcore.EventBus
}

var _ dddcore.Input = (*RegisterUserUseCaseInput)(nil)
var _ dddcore.Output = (*RegisterUserUseCaseOutput)(nil)
var _ dddcore.UseCase[RegisterUserUseCaseInput, RegisterUserUseCaseOutput] = (*RegisterUserUseCase)(nil)

// NewRegisterUserUseCase creates a new instance of RegisterUserUseCase.
func NewRegisterUserUseCase(userRepo repository.UserRepository, idRepo repository.IDRepository, eb dddcore.EventBus) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
		idRepo:   idRepo,
		eventBus: eb,
	}
}

// Execute executes the RegisterUserUseCase with the provided input and returns the output.
func (uc *RegisterUserUseCase) Execute(input *RegisterUserUseCaseInput) (*RegisterUserUseCaseOutput, error) {
	var user entity.User
	var userID int64
	var err error

	if userID, err = uc.idRepo.Incr(1); err != nil {
		return nil, err
	}

	if user, err = entity.NewUser(input.Username, input.Password, userID); err != nil {
		return nil, err
	}

	if err = uc.userRepo.Add(user); err != nil {
		return nil, err
	}

	_ = uc.eventBus.PostAll(user)

	return &RegisterUserUseCaseOutput{
		ID:       user.GetID().String(),
		Username: user.GetUsername(),
		UserID:   user.GetUserID(),
	}, nil
}
