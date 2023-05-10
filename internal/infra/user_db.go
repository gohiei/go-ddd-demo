package infra

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewUserDB() (*gorm.DB, error) {
	dsn := os.Getenv("USER_DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}
