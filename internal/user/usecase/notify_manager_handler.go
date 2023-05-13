package user

import (
	"cypt/internal/dddcore"
	user "cypt/internal/user/entity/events"
	"encoding/json"
	"fmt"
)

// @todo
type NotifyManagerHandler struct {
	eventBus dddcore.EventBus
}

func (h *NotifyManagerHandler) Name() string {
	return "user.notify.manager"
}

func (h *NotifyManagerHandler) EventName() string {
	return user.UserRenamedEventName
}

func (h *NotifyManagerHandler) Handle(eventName string, msg []byte) {
	event := user.UserRenamedEvent{}
	json.Unmarshal(msg, &event)

	fmt.Println("Received: ", event.BaseEvent, event)
}

func NewNotifyManagerHandler(eb dddcore.EventBus) NotifyManagerHandler {
	h := NotifyManagerHandler{}
	eb.Register(&h)

	return h
}
