package db

import "trip-reservation/models"

type UserStore interface {
	GetUserByID(string) (*models.User, error)
}

type MongoUserStore struct{}
