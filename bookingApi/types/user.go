package types

type User struct {
	ID        string `bson:"id,omitempty" json:"id"` // bson - binary json. use ,omitempty to ignore empty fields
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Age       int    `bson:"age" json:"age"`
}
