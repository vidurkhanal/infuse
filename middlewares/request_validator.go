package middlewares

import (
	"fmt"
	"slices"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/vidurkhanal/infuse/constants"
)

type HeaderConfig struct {
	Provider string `json:"provider"`
	Targets  string `json:"targets"`
}

func RequestValidator(c *fiber.Ctx) error {
	if c.Get("content-type") == "" ||
		(c.Get("content-type") != "application/json" &&
			c.Get("content-type") != string(constants.MultipartFormData)) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "either missing content-type headed or invalid content type",
		})
	}

	if (c.Get(fmt.Sprintf("x-%s-mode", constants.POWERED_BY)) == "") ||
		!slices.Contains(constants.Providers,
			c.Get(fmt.Sprintf("x-%s-mode", constants.POWERED_BY))) {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("missing or invalid header: %s", fmt.Sprintf("x-%s-mode", constants.POWERED_BY)),
		})
	}

	if c.Get(fmt.Sprintf("x-%s-config", constants.POWERED_BY)) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("missing header: %s", fmt.Sprintf("x-%s-config", constants.POWERED_BY)),
		})
	}

	var headerConfig HeaderConfig
	if err := json.Unmarshal([]byte(
		c.Get(fmt.Sprintf("x-%s-config",
			constants.POWERED_BY))), &headerConfig); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": fmt.Sprintf("%s must be a valid json",
				fmt.Sprintf("x-%s-config", constants.POWERED_BY)),
		})
	}

	return c.Next()
}
