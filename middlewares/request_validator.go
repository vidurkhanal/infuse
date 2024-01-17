package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vidurkhanal/infuse/constants"
)

func RequestValidator(c *fiber.Ctx) error {
	if c.Get("content-type") == "" ||
		(c.Get("content-type") != "application/json" &&
			c.Get("content-type") != string(constants.MultipartFormData)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "either missing content-type headed or invalid content type",
		})
	}

	if c.Get(fmt.Sprintf("x-%s-mode", constants.POWERED_BY)) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("missing header: %s", fmt.Sprintf("x-%s-mode", constants.POWERED_BY)),
		})
	}

	return c.Next()
}
