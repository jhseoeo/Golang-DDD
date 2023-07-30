package transport

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiberApp() *fiber.App {
	app := fiber.New()
	return app
}
