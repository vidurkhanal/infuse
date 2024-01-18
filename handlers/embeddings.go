package handlers

import "github.com/gofiber/fiber/v2"

func EmbeddingHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Hello": "From Embeddings"})
}
