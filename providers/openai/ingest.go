package openai

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.com/vidurkhanal/infuse/config"
)

func Ingest(c *fiber.Ctx, t config.Target, fn, method, path string) (*fasthttp.Response, error) {
	return &fasthttp.Response{}, nil
}
