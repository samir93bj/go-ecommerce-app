package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Listen("localhost:3000")
}
