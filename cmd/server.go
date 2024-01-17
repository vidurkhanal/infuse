package cmd

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/vidurkhanal/infuse/middlewares"
)

func StartHTTPServer() {
	app := fiber.New(fiber.Config{
		// Marshal and Unmarshal JSON using goccy/go-json, Benchmarked it to be the faster
		// than the standard encoding/json
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Middlewares

	// Idempotency middleware for Fiber allows for fault-tolerant APIs
	// where duplicate requests — for example due to networking issues
	// on the client-side — do not erroneously cause the same action performed
	// multiple times on the server-side.
	//  https://docs.gofiber.io/api/middleware/idempotency
	app.Use(idempotency.New())

	// Monitor middleware for Fiber that
	// reports server metrics,
	// inspired by express-status-monitor
	// https://docs.gofiber.io/api/middleware/monitor
	// Written above the logger middleware so that logger excludes the /metrics endpoint
	app.Get("/metrics", monitor.New(monitor.Config{
		Title:   "Infuse Server Metrics",
		Refresh: time.Duration(1) * time.Second,
	}))

	// Logger middleware for Fiber that logs HTTP request/response details.
	//  https://docs.gofiber.io/api/middleware/logger
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\200b\n",
	}))

	// CORS middleware for Fiber that can be used to enable Cross-Origin
	// Resource Sharing with various options.
	//  https://docs.gofiber.io/api/middleware/cors
	app.Use(cors.New())

	// Liveness and readiness probes middleware for Fiber
	// that provides two endpoints for checking the liveness and readiness
	// state of HTTP applications. Exposes /livez and readyz endpoints.
	//  https://docs.gofiber.io/api/middleware/healthcheck
	app.Use(healthcheck.New())

	// Recover middleware for Fiber that recovers from
	// panics anywhere in the stack chain and handles the
	// control to the centralized ErrorHandler.
	// https://docs.gofiber.io/api/middleware/recover
	app.Use(recover.New())

	app.Use(middlewares.RequestValidator)

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":8080"); err != nil {
		// Graceful shutdown
		app.Shutdown()
	}
}
