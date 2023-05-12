package user

import (
	"cypt/internal/dddcore"
)

type UserRenamedEvent struct {
	*dddcore.BaseEvent
	UserId      string `json:"user_id"`
	OldUsername string `json:"old_username"`
	NewUsername string `json:"new_username"`
}

func NewUserRenameEvent(id, oldUsername, newUsername string) *UserRenamedEvent {
	return &UserRenamedEvent{
		BaseEvent:   dddcore.NewEvent("user.renamed"),
		UserId:      id,
		OldUsername: oldUsername,
		NewUsername: newUsername,
	}
}
