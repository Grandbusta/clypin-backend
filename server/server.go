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

// func (server *Server) initializeRoutes() {
// 	app := server.Router

// 	// http
// 	app.Get("/user/create", http.CreateUser())

// 	// websocket
// 	app.Use("/ws", middlewares.RequestedUpgrade())
// 	app.Get("/ws/send/:user_id", websocket.New(ws.Ws()))

// }

func App() (app *Server) {
	dsn := "host=localhost user=busta password=busta dbname=clypin port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	log.Println("DB connected successfully")
	server := &Server{
		Router: fiber.New(),
		DB:     db,
	}
	return server
}
