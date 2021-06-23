package api

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/parnurzeal/gorequest"
)

// Insert ...
func Insert(message interface{}) error {

	apiUrl := os.Getenv("API_URL")

	fmt.Println("Sending to API")
	// TODO: make a headersconfigurable map
	request := gorequest.New()
	_, responseBody, err := request.Post(apiUrl).Send(message).End()
	if err != nil {
		log.Printf("Error sending to api the message: %v", err)
		return errors.New("Error sending message")
	}
	// TODO: log response properly
	fmt.Println(responseBody)
	return nil
}
