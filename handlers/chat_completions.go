package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vidurkhanal/infuse/utils"
)

func ChatCompletionsHandler(c *fiber.Ctx) error {
	config, err := utils.ConstructConfig(c)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "error",
		})
	}
	log.Printf("config: %#v", config)

	return c.JSON(fiber.Map{"Hello": "World"})
}
