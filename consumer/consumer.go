package consumer

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"

	"cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	r "pubsub-connect/repository"
)

var (
	projectID, subID, gcpCredentials, basuTopic string
	w                                           io.Writer
	dryRun                                      = true
	repositoryConfig                            = "mongo"
	repo                                        r.Repository
)

func init() {

	_err := godotenv.Load()
	if _err != nil {
		fmt.Println("Warning: cant load .env file: " + _err.Error())
	}

	projectID = os.Getenv("GCP_PROJECT_ID")
	subID = os.Getenv("GCP_PUBSUB_SUBSCRIPTION")
	gcpCredentials = os.Getenv("GCP_CREDENTIALS")
	repositoryConfig = os.Getenv("REPOSITORY")
	if os.Getenv("DRY_RUN") == "false" {
		dryRun = false
	}
	fmt.Printf("GCP Poject ID:%s PubSub subscription:%s\n", projectID, subID)
	fmt.Printf("Repository configuration:%s\n", repositoryConfig)

	w = os.Stdout

	switch repositoryConfig {
	case "mongo":
		repo = &r.MongoRepository{}
	case "api":
		repo = &r.ApiRepository{}
	default:
		log.Fatalf("Repository not defined, must be one on of: mongo|api")
	}
}

// Start CIAM events consumer
func Start(fail chan bool) {

	ctx := context.Background()

	var client *pubsub.Client
	var createClientError error

	if gcpCredentials != "" {
		log.Printf("Creating client with credentials: %s", gcpCredentials)
		unescapeGcpCredentials, unescapeError := url.QueryUnescape(gcpCredentials)
		if unescapeError != nil {
			log.Printf("Error unescaping credentials: %v", unescapeError)
		}

		creds, createCredentialsError := google.CredentialsFromJSON(ctx, []byte(unescapeGcpCredentials), secretmanager.DefaultAuthScopes()...)
		if createCredentialsError != nil {
			log.Printf("Error creating credentials from json: %v", createCredentialsError)
		}

		client, createClientError = pubsub.NewClient(ctx, projectID, option.WithCredentials(creds))
		if createClientError != nil {
			log.Printf("Error creating PubSubClient: %v", createClientError)
		}

	} else {
		client, createClientError = pubsub.NewClient(ctx, projectID)
		if createClientError != nil {
			log.Printf("Error creating PubSubClient: %v", createClientError)
		}
	}

	sub := client.Subscription(subID)
	receiveSubError := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message from PubSub: %s", m.Data)

		err := repo.InsertMessage(m.Data)
		if err != nil {
			log.Printf("Cant send message to repository don't ack message from PubSub: %v", err)
		} else {
			if !dryRun {
				m.Ack()
			} else {
				log.Println("Don't ACK dryRun enabled")
			}
		}

	})
	if receiveSubError != nil {
		log.Printf("Error receiving messages from PubSub subscription: %v", receiveSubError)
	}
	fail <- true
}
