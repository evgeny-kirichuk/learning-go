package db

import (
	"context"
	"learning_go/bookingApi/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomCollection = "rooms"

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoRoomStore(client *mongo.Client, dbname string) *MongoRoomStore {
	return &MongoRoomStore{
		client: client,
		coll:   client.Database(dbname).Collection(roomCollection),
	}
}

func (h *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := h.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)
	filter := bson.M{"_id": room.HotelID}
	values := bson.M{
		"$push": bson.M{
			"rooms": room.ID,
		},
	}
	if err := NewMongoHotelStore(h.client, DBname).UpdateHotel(ctx, filter, values); err != nil {
		return nil, err
	}

	return room, nil
}