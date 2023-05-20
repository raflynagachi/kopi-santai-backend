package db

import (
	"fmt"
	"git.garena.com/sea-labs-id/batch-01/rafly-nagachi/final-project-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var db *gorm.DB
var c = config.Config.DBConfig

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		})
}

func Connect() error {
	dsn := config.Config.DatabaseURL
	if config.Config.DatabaseURL == "" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			c.Host, c.User, c.Password, c.DBName, c.Port)
	}
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})
	return err
}

func Get() *gorm.DB {
	return db
}
