package usecase

import (
	"net/http"

	"cypt/internal/dddcore"
	"cypt/internal/user/entity"
	"cypt/internal/user/repository"
)

// RenameUseCaseInput represents the input data for the RenameUseCase.
type RenameUseCaseInput struct {
	ID       string `uri:"id" binding:"required"`
	Username string `form:"username" binding:"required"`
}

// RenameUseCaseOutput represents the output data for the RenameUseCase.
type RenameUseCaseOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// RenameUseCase is a use case for renaming a user.
type RenameUseCase struct {
	userRepo repository.UserRepository
	eventBus dddcore.EventBus
}

// NewRenameUseCase creates a new instance of RenameUseCase.
func NewRenameUseCase(repo repository.UserRepository, eb dddcore.EventBus) *RenameUseCase {
	return &RenameUseCase{
		userRepo: repo,
		eventBus: eb,
	}
}

// Execute executes the RenameUseCase with the provided input and returns the output.
func (uc *RenameUseCase) Execute(input *RenameUseCaseInput) (RenameUseCaseOutput, error) {
	var userID dddcore.UUID
	var user entity.User
	var err error

	if userID, err = dddcore.BuildUUID(input.ID); err != nil || userID.IsNil() {
		return RenameUseCaseOutput{}, dddcore.NewErrorS("10007", "ID is not in UUID format", http.StatusBadRequest)
	}

	if user, err = uc.userRepo.Get(userID); err != nil {
		return RenameUseCaseOutput{}, err
	}

	user.Rename(input.Username)

	if err = uc.userRepo.Rename(user); err != nil {
		return RenameUseCaseOutput{}, err
	}

	uc.eventBus.PostAll(user)

	return RenameUseCaseOutput{
		ID:       user.GetID().String(),
		Username: user.GetUsername(),
	}, nil
}
