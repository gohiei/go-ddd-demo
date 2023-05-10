package user

import (
	"cypt/internal/dddcore"
)

type UserCreatedEvent struct {
	*dddcore.Event
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserCreatedEvent(id, username, password string) UserCreatedEvent {
	return UserCreatedEvent{
		Event:    dddcore.NewEvent("user.created"),
		UserId:   id,
		Username: username,
		Password: password,
	}
}
