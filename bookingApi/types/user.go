package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"` // bson - binary json. use ,omitempty to ignore empty fields
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Age               int                `bson:"age" json:"age"`
	Email             string             `bson:"email" json:"email"`
	EncriptedPassword string             `bson:"encriptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encrpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Age:               params.Age,
		Email:             params.Email,
		EncriptedPassword: string(encrpw),
	}, nil
}
