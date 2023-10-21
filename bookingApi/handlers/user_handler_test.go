package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"learning_go/bookingApi/db"
	"learning_go/bookingApi/types"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testdburi = "mongodb://localhost:27017"
const dbname = "booking_test"

type testdb struct {
	db.UserStore
}

func (tdb *testdb) down(t *testing.T) {
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
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.down(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	apiv1 := app.Group("/api/v1")

	// routes
	apiv1.Post("/user", userHandler.HandleCreateUser)

	params := types.CreateUserParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "ff@ll.com",
		Age:       30,
		Password:  "123qwe123",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/api/v1/user", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	respBody := types.User{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	if respBody.FirstName != params.FirstName {
		t.Errorf("expected %s got %s", params.FirstName, respBody.FirstName)
	}
	if respBody.LastName != params.LastName {
		t.Errorf("expected %s got %s", params.LastName, respBody.LastName)
	}
	if respBody.Email != params.Email {
		t.Errorf("expected %s got %s", params.Email, respBody.Email)
	}
	if respBody.Age != params.Age {
		t.Errorf("expected %d got %d", params.Age, respBody.Age)
	}
}
