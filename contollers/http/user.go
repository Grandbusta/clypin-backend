package http

import (
	"clypin/models"
	"clypin/queries"
	"clypin/utils"

	"github.com/gofiber/fiber/v2"
)

type UserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	var input UserInput
	err := c.BodyParser(&input)
	if err != nil {
		return utils.RespondWIthError(c, fiber.StatusBadRequest, err.Error())
	}
	user := models.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
	}

	new_user := queries.Create(&user)

	return utils.RespondWithJson(c, fiber.StatusCreated, fiber.Map{"data": new_user})
}
