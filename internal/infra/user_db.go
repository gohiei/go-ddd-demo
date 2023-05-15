package infra

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func NewUserDB() (*gorm.DB, error) {
	writeDsn := os.Getenv("USER_WRITE_DB_DSN")
	readDsn := os.Getenv("USER_READ_DB_DSN")

	db, err := gorm.Open(mysql.Open(writeDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return &gorm.DB{}, err
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(readDsn)},
	}))

	return db, nil
}
