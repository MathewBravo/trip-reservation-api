package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bCost              = 12
	firstNameMinLength = 2
	lastNameMinLength  = 2
	minPasswordLength  = 6
)

type CreateUserParams struct {
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	errs := []string{}
	if len(params.FName) < firstNameMinLength {
		errs = append(errs, fmt.Sprintf("first name must be at least %d characters long", firstNameMinLength))
	}
	if len(params.LName) < lastNameMinLength {
		errs = append(errs, fmt.Sprintf("last name must be at least %d characters long", lastNameMinLength))
	}
	if len(params.Password) < minPasswordLength {
		errs = append(errs, fmt.Sprintf("password must be at least %d characters long", minPasswordLength))
	}
	if !validateEmail(params.Email) {
		errs = append(errs, fmt.Sprintf("invalid email"))
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func validateEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e)
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
