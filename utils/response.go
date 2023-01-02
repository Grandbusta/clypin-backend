package utils

import "github.com/gofiber/fiber/v2"

func RespondWIthError(c *fiber.Ctx, code int, message string) error {
	return RespondWithJson(c, code, fiber.Map{"status": code, "error": message})
}

func RespondWithJson(c *fiber.Ctx, code int, payload interface{}) error {
	return c.Status(code).JSON(payload)
}
