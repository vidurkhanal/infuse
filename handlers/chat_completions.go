package handlers

import (
	"log"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/vidurkhanal/infuse/constants"
	"github.com/vidurkhanal/infuse/utils"
)

func ChatCompletionsHandler(c *fiber.Ctx) error {
	t, err := utils.ConstructConfig(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"type":    "error",
		})
	}

	if t.Provider == "" || !slices.Contains(constants.Providers, constants.Provider(t.Provider)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "provider is missing or invalid",
			"type":    "error",
		})
	}

	log.Printf("config: %#v", t)

	return c.JSON(fiber.Map{"Hello": "World"})
}
