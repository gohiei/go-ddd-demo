package infra

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func NewUserDB(writeDsn string, readDsn string) (*gorm.DB, error) {
	// writeDsn := os.Getenv("USER_WRITE_DB_DSN")
	// readDsn := os.Getenv("USER_READ_DB_DSN")
	// writeDsn := viper.GetString("USER_WRITE_DB_DSN")
	// readDsn := viper.GetString("USER_READ_DB_DSN")

	log.Println("new user db")
	log.Println(writeDsn, readDsn)
	log.Println("end new user db")

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
