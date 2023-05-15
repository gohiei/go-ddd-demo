package user

import "cypt/internal/dddcore"

type UserDto struct {
	Id       string
	Username string
}

func NewUserDto(id dddcore.UUID, username string) UserDto {
	return UserDto{
		Id:       id.String(),
		Username: username,
	}
}
