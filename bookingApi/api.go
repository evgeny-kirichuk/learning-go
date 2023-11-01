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



var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func StartServer() {
	addr := flag.String("addr", ":3000", "http service address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	//stores
	userStore := db.NewMongoUserStore(client, db.DBname)
	hotelStore := db.NewMongoHotelStore(client, db.DBname)
	roomStore := db.NewMongoRoomStore(client, db.DBname, hotelStore)

	// handlers initialization
	userHandler := handlers.NewUserHandler(userStore)
	hotelHandler := handlers.NewHotelHandler(roomStore,hotelStore)

	// routes
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Post("/user", userHandler.HandleCreateUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Put("/user/:id", userHandler.HandleUpdateUser)
	apiv1.Get("/user/:id", userHandler.HandleGetUserById)

	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)

	app.Listen(*addr)
}
