package main

import (
	"clypin/config"
	"clypin/contollers/http"
	"clypin/contollers/ws"
	"clypin/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func init() {
	//Load env files
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// } else {
	// 	fmt.Println(".env loaded")
	// }
	config.ConnectDB()
}

func handleRequests() {
	app := fiber.New()

	// User Routes
	user := app.Group("/user")
	user.Post("/create", http.CreateUser)

	// websocket
	app.Use("/ws", middlewares.RequestedUpgrade())
	app.Get("/ws/send/:user_id", websocket.New(ws.Ws()))

	log.Fatal(app.Listen(":3000"))
}

func main() {
	handleRequests()
}
