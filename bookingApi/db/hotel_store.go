package db

import (
	"context"
	"learning_go/bookingApi/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const hotelCollection = "hotels"

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(ctx context.Context, filter bson.M, values bson.M) error
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbname).Collection(hotelCollection),
	}
}

func (h *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := h.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}


func (s *MongoHotelStore) UpdateHotel(ctx context.Context, filter bson.M, values bson.M) error {
	// update := bson.M{
	// 	"$set": values,
	// }
	_, err := s.coll.UpdateOne(ctx, filter, values)
	if err != nil {
		return err
	}
	return nil
}