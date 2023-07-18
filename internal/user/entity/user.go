// Package entity defines the user entity and related operations.
package entity

import (
	"net/http"

	"cypt/internal/dddcore"
	"cypt/internal/user/entity/events"
)

// User represents a user entity.
type User struct {
	*dddcore.BaseAggregateRoot
	id       dddcore.UUID
	username string
	password string
	userID   int64
}

// NewUser creates a new User instance with the provided username, password, and userID.
func NewUser(username string, password string, userID int64) (User, error) {
	if username == "" {
		return User{}, dddcore.NewErrorS("10001", "missing value `username`", http.StatusBadRequest)
	}

	user := User{
		BaseAggregateRoot: dddcore.NewAggregateRoot(),
		id:                dddcore.NewUUID(),
		username:          username,
		password:          password,
		userID:            userID,
	}

	user.AddDomainEvent(events.NewUserCreatedEvent(
		user.id.String(),
		user.username,
		user.password,
		user.userID,
	))

	return user, nil
}

// BuildUser creates a User instance with the provided ID, username, password, and userID.
func BuildUser(id string, username string, password string, userID int64) User {
	uid, _ := dddcore.BuildUUID(id)

	return User{
		BaseAggregateRoot: dddcore.NewAggregateRoot(),
		id:                uid,
		username:          username,
		password:          password,
		userID:            userID,
	}
}

// GetID returns the ID of the user.
func (u *User) GetID() dddcore.UUID {
	return u.id
}

// SetID sets the ID of the user.
func (u *User) SetID(id dddcore.UUID) {
	u.id = id
}

// GetUsername returns the username of the user.
func (u *User) GetUsername() string {
	return u.username
}

// GetPassword returns the password of the user.
func (u *User) GetPassword() string {
	return u.password
}

// GetUserID returns the userID of the user.
func (u *User) GetUserID() int64 {
	return u.userID
}

// Rename renames the user with the provided username.
func (u *User) Rename(username string) {
	old := u.username

	// if old == username {
	// 	return
	// }

	u.username = username

	u.AddDomainEvent(events.NewUserRenameEvent(
		u.id.String(),
		old,
		username,
	))
}
