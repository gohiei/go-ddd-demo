package user

import "time"

type UserModel struct {
	ID        string `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (UserModel) TableName() string {
	return "user"
}
