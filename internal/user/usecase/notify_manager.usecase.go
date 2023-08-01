package usecase

import (
	"encoding/json"
	"fmt"

	"cypt/internal/dddcore"
	"cypt/internal/user/entity/events"
	"cypt/internal/user/repository"
)

var _ dddcore.UseCase[NotifyManagerUseCaseInput, NotifyManagerUseCaseOutput] = (*NotifyManagerUseCase)(nil)

// NotifyManagerUseCase is a handler for NotifyManager events.
type NotifyManagerUseCase struct {
	repo repository.OutsideRepository
}

func (h *NotifyManagerUseCase) Name() string {
	return "user.notify.manager"
}

func (h *NotifyManagerUseCase) EventName() string {
	return events.UserRenamedEventName
}

func (h *NotifyManagerUseCase) When(eventName string, msg []byte) {
	event := events.UserRenamedEvent{}

	if err := json.Unmarshal(msg, &event); err != nil {
		return
	}

	_, _ = h.Execute(&NotifyManagerUseCaseInput{})
}

type NotifyManagerUseCaseInput struct{}
type NotifyManagerUseCaseOutput struct{ Result bool }

func (h *NotifyManagerUseCase) Execute(input *NotifyManagerUseCaseInput) (*NotifyManagerUseCaseOutput, error) {
	if data, err := h.repo.GetEchoData(); err != nil {
		// nolint: forbidigo
		fmt.Println("Echo: ", data)
	}

	return &NotifyManagerUseCaseOutput{Result: false}, nil
}

func NewNotifyManagerHandler(repo repository.OutsideRepository, eb dddcore.EventBus) NotifyManagerUseCase {
	h := NotifyManagerUseCase{repo: repo}
	eb.Register(&h)

	return h
}
