package repository

import (
	"fmt"
	"log"
	"os"

	m "pubsub-connect/repository/mongo"

	"github.com/joho/godotenv"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

var (
	apiUrl     string
	apiType    string
	apiCountry string
)

func init() {

	_err := godotenv.Load()
	if _err != nil {
		fmt.Println("Warning: cant load .env file: " + _err.Error())
	}

	apiUrl = os.Getenv("API_URL")
}

// Repository ...
type Repository interface {
	InsertMessage(message interface{}) error
}

// MongoRepository ...
type MongoRepository struct{}

// InsertMessage in MongoRepository
func (r *MongoRepository) InsertMessage(message interface{}) error {
	// TODO: Validate message with json format
	fmt.Println("Insert to MongoRepository")
	err := m.Insert(message)
	if err != nil {
		log.Printf("Cant save message to Mongo: %v", err)
		return errors.New("")
	}
	return nil
}

// BasuRepository ...
type ApiRepository struct{}

// InsertMessage in BasuRepository
func (r *ApiRepository) InsertMessage(message interface{}) error {
	fmt.Println("Send to API")
	// TODO: make headers a configurable map
	request := gorequest.New()
	_, responseBody, err := request.Post(apiUrl).Send(message).End()
	if err != nil {
		log.Printf("Error send api message: %v", err)
		return errors.New("Error creating message")
	}
	fmt.Println("go api")
	fmt.Println(responseBody)
	return nil
}
