package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Router *fiber.App
	DB     *gorm.DB
}

var DB *gorm.DB

func App() (app *Server) {
	dsn := "host=localhost user=busta password=busta dbname=clypin port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	log.Println("DB connected successfully")
	RunMigrations(db)
	server := &Server{
		Router: fiber.New(),
		DB:     db,
	}
	DB = db

	return server
}
