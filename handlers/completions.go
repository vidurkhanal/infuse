package handlers

import "github.com/gofiber/fiber/v2"

func CompletionsHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Hello": "From Completions"})
}
