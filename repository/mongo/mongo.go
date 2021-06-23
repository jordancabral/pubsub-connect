package mongo

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client                                                   *mongo.Client
	mongoURL, mongoUser, mongoPass, mongoDB, mongoCollection string
)

func init() {

	// TODO: No esta bueno que se intancie solo por la config de env.
	if os.Getenv("REPOSITORY") == "mongo" {
		mongoURL = os.Getenv("MONGO_DB_URI")
		mongoUser = os.Getenv("MONGO_DB_USER")
		mongoPass = os.Getenv("MONGO_DB_PASS")
		mongoDB = os.Getenv("MONGO_DB")
		mongoCollection = os.Getenv("MONGO_COLLECTION")

		var err error
		mongoOptions := options.Client().ApplyURI(mongoURL)
		if mongoUser != "" {
			credentials := options.Credential{
				Username: mongoUser,
				Password: mongoPass,
			}
			mongoOptions.SetAuth(credentials)
		}
		client, err = mongo.NewClient(mongoOptions)
		if err != nil {
			log.Fatal(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		//	defer client.Disconnect(ctx)
	}
}

// Insert ...
func Insert(message interface{}) error {
	ctx := context.Background()

	collection := client.Database(mongoDB).Collection(mongoCollection)

	res, err := collection.InsertOne(ctx, message)
	if err != nil {
		log.Printf("Mongo error: %v", err)
		return errors.New("")
	}
	id := res.InsertedID.(primitive.ObjectID)
	log.Printf("Inserted document with id:%s", id.String())
	return nil
}
