package user

import (
	dddcore "cypt/internal/dddcore"
	entity "cypt/internal/user/entity"
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

func NewRenameUseCase(repo repo.UserRepository, eb dddcore.EventBus) *RenameUseCase {
	return &RenameUseCase{
		userRepo: repo,
		eventBus: eb,
	}
}

func (uc RenameUseCase) Execute(input *RenameUseCaseInput) (RenameUseCaseOutput, error) {
	var user entity.User
	var err error

	if user, err = uc.userRepo.Get(input.ID); err != nil {
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
