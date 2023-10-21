package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func ToObjectID(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return oid
}
