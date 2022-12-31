package main

import (
	"clypin/contollers/http"
	"clypin/contollers/ws"
	"clypin/middlewares"
	"clypin/server"

	"github.com/gofiber/websocket/v2"
)

func main() {
	app := server.App().Router
	// InitializeRoutes(s)

	app.Get("/user/create", http.CreateUser())

	// websocket
	app.Use("/ws", middlewares.RequestedUpgrade())
	app.Get("/ws/send/:user_id", websocket.New(ws.Ws()))

	app.Listen(":3000")
}
