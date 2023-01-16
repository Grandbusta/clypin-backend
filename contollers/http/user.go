package http

import (
	"clypin/models"
	"clypin/queries"
	"clypin/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	_, err = queries.FindUserByEmail(input.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.RespondWIthError(c, fiber.ErrConflict.Code, "user already registered")
	}

	new_user, err := queries.CreateUser(&user)
	if err != nil {
		return utils.RespondWIthError(c,
			fiber.StatusInternalServerError,
			"An error occured",
		)
	}

	return utils.RespondWithJson(c,
		fiber.StatusCreated,
		fiber.Map{"data": new_user},
	)
}

func LoginUser(c *fiber.Ctx) error {
	var input LoginUserInput
	err := c.BodyParser(&input)
	if err != nil {
		return utils.RespondWIthError(c, fiber.StatusBadRequest, err.Error())
	}
	res, err := queries.FindUserByEmail(input.Email)
	if err != nil {
		panic(err.Error())
	}

	return utils.RespondWithJson(c, fiber.StatusOK, res)
}
