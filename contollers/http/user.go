package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateUser() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		log.Println("Getting to create user")
		return c.SendString("Hello, World 👋!")
	}
}
