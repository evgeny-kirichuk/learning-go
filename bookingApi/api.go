package bookingApi

import (
	"context"
	"flag"
	"learning_go/bookingApi/db"
	"learning_go/bookingApi/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func StartServer() {
	addr := flag.String("addr", ":3000", "http service address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	// handlers initialization
	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))

	// routes
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Post("/user", userHandler.HandleCreateUser)
	apiv1.Get("/user/:id", userHandler.HandleGetUserById)

	app.Listen(*addr)
}
