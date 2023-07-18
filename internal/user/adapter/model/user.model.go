// Package model provides the data model for the user entity in the database.
package model

import "time"

// UserModel represents the user model in the database.
type UserModel struct {
	ID        string `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int64 `gorm:"column:user_id"`
}

// TableName returns the name of the database table for the UserModel.
func (UserModel) TableName() string {
	return "user"
}
