package main

import (
	"clypin/config"
	"clypin/contollers/http"
	"clypin/contollers/ws"
	"clypin/middlewares"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
)

func init() {
	// Load env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println(".env loaded")
	}
	config.ConnectDB()
}

func handleRequests() {
	app := fiber.New()

	// User Routes
	user := app.Group("/user")
	user.Post("/create", http.CreateUser)
	user.Post("/login", http.LoginUser)

	// websocket
	app.Use("/ws", middlewares.RequestedUpgrade())
	app.Get("/ws/send/:user_id", websocket.New(ws.Ws()))
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3000"
	}
	log.Fatal(app.Listen(PORT))
}

func main() {
	handleRequests()
}
