package user

import (
	"errors"

	"cypt/internal/dddcore"
	event "cypt/internal/user/entity/events"
)

var (
	ErrUserMissingValues = errors.New("missing values: username")
)

type User struct {
	dddcore.AggregateRoot
	id       dddcore.UUID
	username string
	password string
}

func NewUser(username string, password string) (User, error) {
	if username == "" {
		return User{}, ErrUserMissingValues
	}

	user := User{
		id:       dddcore.NewUUID(),
		username: username,
		password: password,
	}

	user.AddDomainEvent(event.NewUserCreatedEvent(
		user.id.String(),
		user.username,
		user.password,
	))

	return user, nil
}

func BuildUser(id string, username string, password string) User {
	uid, _ := dddcore.BuildUUID(id)

	return User{
		id:       uid,
		username: username,
		password: password,
	}
}

func (u *User) GetId() dddcore.UUID {
	return u.id
}

func (u *User) SetId(id dddcore.UUID) {
	u.id = id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) Rename(username string) {
	old := u.username
	u.username = username

	u.AddDomainEvent(event.NewUserRenameEvent(
		u.id.String(),
		old,
		username,
	))
}
