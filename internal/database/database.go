package database

import (
	"fmt"
	"log"
	"os"
	"pet-matching-service/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db

	DB.AutoMigrate(&model.Pet{})
}

func Close() {
	DB.Close()
}
