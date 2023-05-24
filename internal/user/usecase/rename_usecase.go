package user

import (
	dddcore "cypt/internal/dddcore"
	repo "cypt/internal/user/repository"
)

type RenameUseCaseInput struct {
	ID       dddcore.UUID `uri:"id" binding:"required"`
	Username string       `form:"username" binding:"required"`
}

type RenameUseCaseOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewRenameUseCaseInput(id string, username string) RenameUseCaseInput {
	uuid, _ := dddcore.BuildUUID(id)

	return RenameUseCaseInput{ID: uuid, Username: username}
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
	user, err := uc.userRepo.Get(input.ID)

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
		ID:       user.GetID().String(),
		Username: user.GetUsername(),
	}, nil
}
