package bookingApi

import "github.com/gofiber/fiber"

func StartServer() {
	app := fiber.New()
	app.Get("/bookings", handleBookings)
	app.Listen(":3000")
}

func handleBookings(c *fiber.Ctx) {
	c.JSON(map[string]string{"booking": "booked"})
}