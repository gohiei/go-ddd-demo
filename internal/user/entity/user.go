package user

import (
	"cypt/internal/dddcore"
	event "cypt/internal/user/entity/events"
	exception "cypt/internal/user/exception"
)

type User struct {
	*dddcore.BaseAggregateRoot
	id       dddcore.UUID
	username string
	password string
}

func NewUser(username string, password string) (User, error) {
	if username == "" {
		return User{}, exception.NewErrMissingValueUsername()
	}

	user := User{
		BaseAggregateRoot: dddcore.NewAggregateRoot(),
		id:                dddcore.NewUUID(),
		username:          username,
		password:          password,
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
		BaseAggregateRoot: dddcore.NewAggregateRoot(),
		id:                uid,
		username:          username,
		password:          password,
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
