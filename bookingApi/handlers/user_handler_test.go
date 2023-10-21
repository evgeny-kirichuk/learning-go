package handlers

import (
	"context"
	"learning_go/bookingApi/db"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testdburi = "mongodb://localhost:27017"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) down(t *testing.T){
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.down(t)
}
