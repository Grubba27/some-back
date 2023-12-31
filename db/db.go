package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Error connecting to database")
		log.Panic(err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		return InitDB()
	}
	return DB
}
