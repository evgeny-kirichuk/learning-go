package bookingApi

import "github.com/gofiber/fiber"

func StartServer() {
	app := fiber.New()
	app.Listen(":3000")
}