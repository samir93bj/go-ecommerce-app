package api

import (
	config "go-ecommerce-app/configs"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}
