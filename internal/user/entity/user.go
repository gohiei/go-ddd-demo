package user

import (
	"net/http"

	"cypt/internal/dddcore"
	event "cypt/internal/user/entity/events"
)

type User struct {
	*dddcore.BaseAggregateRoot
	id       dddcore.UUID
	username string
	password string
	userID   int64
}

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

	user.AddDomainEvent(event.NewUserCreatedEvent(
		user.id.String(),
		user.username,
		user.password,
		user.userID,
	))

	return user, nil
}

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

func (u *User) GetID() dddcore.UUID {
	return u.id
}

func (u *User) SetID(id dddcore.UUID) {
	u.id = id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetUserID() int64 {
	return u.userID
}

func (u *User) Rename(username string) {
	old := u.username

	// if old == username {
	// 	return
	// }

	u.username = username

	u.AddDomainEvent(event.NewUserRenameEvent(
		u.id.String(),
		old,
		username,
	))
}
