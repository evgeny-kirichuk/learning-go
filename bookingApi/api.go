package bookingApi

import (
	"flag"
	"learning_go/bookingApi/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	addr := flag.String("addr", ":3000", "http service address")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", handlers.HandleGetUsers)
	apiv1.Get("/user/:id", handlers.HandleGetUserById)

	app.Listen(*addr)
}
