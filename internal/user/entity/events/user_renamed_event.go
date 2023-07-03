package user

import "cypt/internal/dddcore"

// UserRenamedEventName represents the event name
const (
	UserRenamedEventName = "user.renamed"
)

// UserRenamedEvent represents the event when a user is renamed.
type UserRenamedEvent struct {
	*dddcore.BaseEvent
	UserID      string `json:"user_id"`
	OldUsername string `json:"old_username"`
	NewUsername string `json:"new_username"`
}

// NewUserRenameEvent creates a new UserRenamedEvent instance with the provided parameters.
func NewUserRenameEvent(id, oldUsername, newUsername string) *UserRenamedEvent {
	return &UserRenamedEvent{
		BaseEvent:   dddcore.NewEvent(UserRenamedEventName),
		UserID:      id,
		OldUsername: oldUsername,
		NewUsername: newUsername,
	}
}
