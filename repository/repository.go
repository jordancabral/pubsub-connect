package repository

import (
	"fmt"
	"pubsub-connect/repository/api"
	"pubsub-connect/repository/mongo"

	"github.com/pkg/errors"
)

// Repository ...
type Repository interface {
	InsertMessage(message interface{}) error
}

// MongoRepository ...
type MongoRepository struct{}

// InsertMessage in MongoRepository
func (r *MongoRepository) InsertMessage(message interface{}) error {
	// TODO: Validate message with json format
	fmt.Println("Inserting to MongoRepository")
	err := mongo.Insert(message)
	if err != nil {
		return errors.New("Error inserting message")
	}
	return nil
}

// ApiRepository ...
type ApiRepository struct{}

// InsertMessage in ApiRepository
func (r *ApiRepository) InsertMessage(message interface{}) error {
	fmt.Println("Sending to API")
	err := api.Insert(message)
	if err != nil {
		return errors.New("Error inserting message")
	}
	return nil
}
