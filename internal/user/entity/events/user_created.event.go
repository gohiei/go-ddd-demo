// Package events defines the domain events related to the user entity.
package events

import "cypt/internal/dddcore"

// UserCreatedEventName represents the event name
const (
	UserCreatedEventName = "user.created"
)

// UserCreatedEvent represents the event when a user is created.
type UserCreatedEvent struct {
	*dddcore.BaseEvent
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	UserIntID int64  `json:"user_int_id"`
}

var _ dddcore.Event = (*UserCreatedEvent)(nil)

// NewUserCreatedEvent creates a new UserCreatedEvent instance with the provided parameters.
func NewUserCreatedEvent(id, username, password string, userID int64) *UserCreatedEvent {
	return &UserCreatedEvent{
		BaseEvent: dddcore.NewEvent(UserCreatedEventName),
		UserID:    id,
		Username:  username,
		Password:  password,
		UserIntID: userID,
	}
}
