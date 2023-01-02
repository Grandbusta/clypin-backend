package http

import (
	"clypin/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateUser() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := models.User{
			Email:     "bustajay30@gmail.com",
			FirstName: "Bolu",
			LastName:  "Busta",
			Password:  "Mil",
		}

		fmt.Println(models.Create(user))

		log.Println("Getting to create user")
		return c.SendString("Hello, World 👋!")
	}
}
