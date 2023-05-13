package user

import (
	"cypt/internal/dddcore"
)

const (
	UserCreatedEventName = "user.created"
)

type UserCreatedEvent struct {
	*dddcore.BaseEvent
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var _ dddcore.Event = (*UserCreatedEvent)(nil)

func NewUserCreatedEvent(id, username, password string) *UserCreatedEvent {
	return &UserCreatedEvent{
		BaseEvent: dddcore.NewEvent(UserCreatedEventName),
		UserId:    id,
		Username:  username,
		Password:  password,
	}
}
