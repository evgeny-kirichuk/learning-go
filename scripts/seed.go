package main

import (
	"context"
	"fmt"
	"learning_go/bookingApi/db"
	"learning_go/bookingApi/types"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())


	hotelStore := db.NewMongoHotelStore(client, db.DBname)
	roomStore := db.NewMongoRoomStore(client, db.DBname, hotelStore)

	hotel := types.Hotel{
		Name:    "Hilton",
		Address: "1234 Main St",
		Rooms: []primitive.ObjectID{},
	}

	room := types.Room{
		Type:      types.Single,
		BasePrice: 100,
		Price:     100,
	}

	indertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room.HotelID = indertedHotel.ID

	insertedRoom, err := roomStore.InsertRoom(ctx, &room)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(indertedHotel, insertedRoom)
}
