package main

import (
	"clypin/contollers/http"
	"clypin/contollers/ws"
	"clypin/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Server struct {
	Router *fiber.App
}

func InitializeRoutes(server *Server) {
	app := server.Router

	// http
	app.Get("/user/create", http.CreateUser())

	// websocket
	app.Use("/ws", middlewares.RequestedUpgrade())
	app.Get("/ws/send/:user_id", websocket.New(ws.Ws()))

}

func main() {
	app := fiber.New()
	server := &Server{
		Router: app,
	}
	InitializeRoutes(server)

	app.Listen(":3000")
}
