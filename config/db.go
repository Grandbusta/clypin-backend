package config

import (
	"clypin/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", db_host, db_user, db_pass, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	if err != nil {
		log.Fatal("error connecting to DB")
	} else {
		log.Println("DB conncted successfully")
	}
	// Migrate
	db.AutoMigrate(
		&models.User{},
		&models.Message{},
	)

	DB = db
}
