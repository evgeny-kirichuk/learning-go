package db

import (
	"context"
	"fmt"
	"learning_go/bookingApi/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, values bson.M) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(userCollection),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("dropping users collection")
	return s.coll.Drop(ctx)
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	curr, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := curr.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, values bson.M) error {
	update := bson.M{
		"$set": values,
	}
	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}