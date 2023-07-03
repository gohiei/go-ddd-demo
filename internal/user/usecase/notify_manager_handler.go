package user

import (
	"encoding/json"
	"fmt"

	"cypt/internal/dddcore"
	user "cypt/internal/user/entity/events"
)

// NotifyManagerHandler is a handler for NotifyManager events.
type NotifyManagerHandler struct {
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
}

// NewNotifyManagerHandler creates a new instance of NotifyManagerHandler and registers it to the event bus.
func NewNotifyManagerHandler(eb dddcore.EventBus) NotifyManagerHandler {
	h := NotifyManagerHandler{}
	eb.Register(&h)

	return h
}
