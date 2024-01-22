package handlers

import "github.com/gofiber/fiber/v2"

func ChatCompletionsHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Hello": "World"})
}
