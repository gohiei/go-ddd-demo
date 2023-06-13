package user

import (
	"encoding/json"
	"fmt"

	"cypt/internal/dddcore"
	user "cypt/internal/user/entity/events"
)

// @todo
type NotifyManagerHandler struct {
}

func (h *NotifyManagerHandler) Name() string {
	return "user.notify.manager"
}

func (h *NotifyManagerHandler) EventName() string {
	return user.UserRenamedEventName
}

func (h *NotifyManagerHandler) When(eventName string, msg []byte) {
	event := user.UserRenamedEvent{}
	json.Unmarshal(msg, &event)

	fmt.Println("Received: ", event.BaseEvent, event)
}

func NewNotifyManagerHandler(eb dddcore.EventBus) NotifyManagerHandler {
	h := NotifyManagerHandler{}
	eb.Register(&h)

	return h
}
