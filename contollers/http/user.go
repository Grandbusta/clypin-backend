package http

import (
	"clypin/models"
	"clypin/queries"
	"clypin/utils"

	"github.com/gofiber/fiber/v2"
)

type CreateUserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *fiber.Ctx) error {
	var input CreateUserInput
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

func LoginUser(c *fiber.Ctx) error {
	var input LoginUserInput
	err := c.BodyParser(&input)
	if err != nil {
		return utils.RespondWIthError(c, fiber.StatusBadRequest, err.Error())
	}
	res, err := queries.FindByEmail(input.Email)
	if err != nil {
		panic(err.Error())
	}
	return utils.RespondWithJson(c, fiber.StatusOK, res)
}
