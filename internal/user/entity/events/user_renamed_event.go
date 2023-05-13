package user

import (
	"cypt/internal/dddcore"
)

const (
	UserRenamedEventName = "user.renamed"
)

type UserRenamedEvent struct {
	*dddcore.BaseEvent
	UserId      string `json:"user_id"`
	OldUsername string `json:"old_username"`
	NewUsername string `json:"new_username"`
}

func NewUserRenameEvent(id, oldUsername, newUsername string) *UserRenamedEvent {
	return &UserRenamedEvent{
		BaseEvent:   dddcore.NewEvent(UserRenamedEventName),
		UserId:      id,
		OldUsername: oldUsername,
		NewUsername: newUsername,
	}
}
