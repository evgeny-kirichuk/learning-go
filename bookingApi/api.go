package bookingApi

import (
	"context"
	"flag"
	"fmt"
	"learning_go/bookingApi/handlers"
	"learning_go/bookingApi/types"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "booking"
const userCollection = "users"

func StartServer() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	userColl := client.Database(dbname).Collection(userCollection)

	testUser := types.User{
		FirstName: "Evgeny",
		LastName:  "Kirichuk",
		Age:       30,
	}
	res, err := userColl.InsertOne(context.Background(), testUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %+v", res)
	addr := flag.String("addr", ":3000", "http service address")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", handlers.HandleGetUsers)
	apiv1.Get("/user/:id", handlers.HandleGetUserById)

	app.Listen(*addr)
}
