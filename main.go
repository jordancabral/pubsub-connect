package main

import (
	"fmt"
	"net/http"

	"pubsub-connect/consumer"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting...")

	_err := godotenv.Load()
	if _err != nil {
		fmt.Println("Warning: cant load .env file: " + _err.Error())
	}

	// Start consumer
	consumerFail := false
	fail := make(chan bool)
	go consumer.Start(fail)
	go func() {
		consumerFail = <-fail
	}()

	// Healthcheck
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		if consumerFail {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Consumer Error."))
		} else {
			fmt.Fprintf(w, "Up and running.\n")
		}
	})

	http.ListenAndServe(":8080", nil)
}
