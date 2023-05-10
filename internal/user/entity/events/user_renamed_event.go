package user

import (
	"cypt/internal/dddcore"
)

type UserRenamedEvent struct {
	*dddcore.Event
	UserId      string `json:"user_id"`
	OldUsername string `json:"old_username"`
	NewUsername string `json:"new_username"`
}

func NewUserRenameEvent(id, oldUsername, newUsername string) UserRenamedEvent {
	return UserRenamedEvent{
		Event:       dddcore.NewEvent("user.renamed"),
		UserId:      id,
		OldUsername: oldUsername,
		NewUsername: newUsername,
	}
}
