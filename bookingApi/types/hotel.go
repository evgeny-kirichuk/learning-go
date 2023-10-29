package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID      primitive.ObjectID   `bson:"id,omitempty" json:"id,omitempty"` // bson - binary
	Name    string               `bson:"name" json:"name"`
	Address string               `bson:"address" json:"address"`
	Rooms   []primitive.ObjectID `bson:"rooms" json:"rooms"`
}

type RoomType int

const (
	_ RoomType = iota
	Single
	Double
)

type Room struct {
	ID        primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	Type      RoomType           `bson:"type" json:"type"`
	BasePrice float64            `bson:"basePrice" json:"basePrice"`
	Price     float64            `bson:"price" json:"price"`
	HotelID   primitive.ObjectID `bson:"hotelId" json:"hotelId"`
}
