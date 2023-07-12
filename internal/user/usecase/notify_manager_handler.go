package user

import (
	"encoding/json"
	"fmt"

	"cypt/internal/dddcore"
	user "cypt/internal/user/entity/events"
	repository "cypt/internal/user/repository"
)

// NotifyManagerHandler is a handler for NotifyManager events.
type NotifyManagerHandler struct {
	repo repository.OutsideRepository
}

// Name returns the name of the handler.
func (h *NotifyManagerHandler) Name() string {
	return "user.notify.manager"
}

// EventName returns the name of the event handled by the handler.
func (h *NotifyManagerHandler) EventName() string {
	return user.UserRenamedEventName
}

// When is the actual processing logic of the event handler, used to handle specific events.
func (h *NotifyManagerHandler) When(eventName string, msg []byte) {
	event := user.UserRenamedEvent{}
	json.Unmarshal(msg, &event)

	fmt.Println("NotifyManagerHandler Received:", event.BaseEvent, event)

	if data, err := h.repo.GetEchoData(); err == nil {
		fmt.Println("Echo: ", data)
	}
}

// NewNotifyManagerHandler creates a new instance of NotifyManagerHandler and registers it to the event bus.
func NewNotifyManagerHandler(repo repository.OutsideRepository, eb dddcore.EventBus) NotifyManagerHandler {
	h := NotifyManagerHandler{repo: repo}
	eb.Register(&h)

	return h
}
