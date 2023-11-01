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
	Insert(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(ctx context.Context, filter bson.M, values bson.M) error
	List(ctx context.Context, filter bson.M) ([]*types.Hotel, error)
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



func (h *MongoHotelStore) Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := h.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}


func (s *MongoHotelStore) Update(ctx context.Context, filter bson.M, values bson.M) error {
	_, err := s.coll.UpdateOne(ctx, filter, values)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoHotelStore) List(ctx context.Context, filter bson.M) ([]*types.Hotel, error) {
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var hotels []*types.Hotel
	if err := cur.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}