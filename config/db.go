package config

import (
	"clypin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=busta password=busta dbname=clypin port=5432"
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
	db.Debug().AutoMigrate(
		&models.User{},
		&models.Message{},
	)

	DB = db
}
