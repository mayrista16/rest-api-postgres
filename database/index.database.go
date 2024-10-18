package database

import (
	"fmt"
	"log"

	"github.com/mayrista16/rest-api-postgres/configs/app_config/db_config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var errConnection error

	if db_config.DB_DRIVER == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", db_config.DB_HOST, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME, db_config.DB_PORT)
		DB, errConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if errConnection != nil {
		panic("Database not connected !")
	}

	log.Println("Connected to database")
}
