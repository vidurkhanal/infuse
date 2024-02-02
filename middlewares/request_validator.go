package middlewares

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/vidurkhanal/infuse/constants"
)

func RequestValidator(c *fiber.Ctx) error {
	if c.Get(fiber.HeaderContentType) == "" ||
		(c.Get(fiber.HeaderContentType) != fiber.MIMEApplicationJSON &&
			c.Get(fiber.HeaderContentType) != fiber.MIMEMultipartForm) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "either missing content-type headed or invalid content type",
		})
	}

	if (c.Get(fmt.Sprintf("x-%s-mode", constants.POWERED_BY)) == "") &&
		(c.Get(fmt.Sprintf("x-%s-provider", constants.POWERED_BY)) == "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": fmt.Sprintf("missing or invalid header: %s or %s",
				fmt.Sprintf("x-%s-mode", constants.POWERED_BY),
				fmt.Sprintf("x-%s-provider", constants.POWERED_BY)),
		})
	}

	if c.Get(fmt.Sprintf("x-%s-config", constants.POWERED_BY)) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": fmt.Sprintf("missing header: %s ",
				fmt.Sprintf("x-%s-config", constants.POWERED_BY),
			),
		})
	}

	var headerConfig HeaderConfig
	if err := json.Unmarshal([]byte(
		c.Get(fmt.Sprintf("x-%s-config",
			constants.POWERED_BY))), &headerConfig); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": fmt.Sprintf("%s must be a valid json \n %s",
				fmt.Sprintf("x-%s-config", constants.POWERED_BY), err),
		})
	}

	if err := headerConfig.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Next()
}
