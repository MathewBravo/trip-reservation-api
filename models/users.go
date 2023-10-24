package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bCost = 12
)

type CreateUserParams struct {
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FName       string             `bson:"f_name" json:"f_name"`
	LName       string             `bson:"l_name" json:"l_name"`
	Email       string             `bson:"email" json:"email"`
	EncPassword string             `bson:"enc_password" json:"-"`
}

func CreateUserFromParams(params CreateUserParams) (*User, error) {
	enps, err := bcrypt.GenerateFromPassword([]byte(params.Password), bCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FName:       params.FName,
		LName:       params.LName,
		Email:       params.Email,
		EncPassword: string(enps),
	}, nil
}
