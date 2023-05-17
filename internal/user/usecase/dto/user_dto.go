package user

import "cypt/internal/dddcore"

type UserDto struct {
	ID       string
	Username string
}

func NewUserDto(id dddcore.UUID, username string) UserDto {
	return UserDto{
		ID:       id.String(),
		Username: username,
	}
}
