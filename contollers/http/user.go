package http

import (
	"clypin/models"
	"clypin/queries"
	"clypin/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := models.User{
		Email:     "bustajay30@gmail.com",
		FirstName: "Bolu",
		LastName:  "Busta",
		Password:  "Mil",
	}

	new_user := queries.Create(&user)

	return utils.RespondWithJson(c, fiber.StatusCreated, new_user)
}
