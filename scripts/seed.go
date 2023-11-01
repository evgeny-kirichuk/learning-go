package main

import (
	"context"
	"learning_go/bookingApi/db"
	"learning_go/bookingApi/types"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	roomStore db.RoomStore
	hotelStore db.HotelStore
	ctx = context.Background()
)

func seedHotel(name, address string) {
	hotel := types.Hotel{
		Name:    name,
		Address: address,
		Rooms: []primitive.ObjectID{},
	}

	rooms := []types.Room{
		{
			Type:      types.Single,
			BasePrice: 100,
			Price:     100,
		},
		{
			Type:      types.Double,
			BasePrice: 200,
			Price:     200,
		},
		{
			Type:      types.Single,
			BasePrice: 300,
			Price:     300,
		},
	}

		insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
		if err != nil {
			log.Fatal(err)
		}
		for _, room := range rooms {
			room.HotelID = insertedHotel.ID
			insertedRoom, err := roomStore.InsertRoom(ctx, &room)
			if err != nil {
				log.Fatal(err)
			}
			insertedHotel.Rooms = append(insertedHotel.Rooms, insertedRoom.ID)
		}
}

func main() {
	defer client.Disconnect(context.Background())
	seedHotel("Hilton", "1234 Main Street")
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBname).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client, db.DBname)
	roomStore = db.NewMongoRoomStore(client, db.DBname, hotelStore)
}