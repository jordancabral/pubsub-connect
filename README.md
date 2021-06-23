# pubsub connect #

This service read messages from PubSub and send them to different configurable repositories (mongo, custom API)

## Config ##

GCP_PROJECT_ID
GCP_PUBSUB_SUBSCRIPTION
GCP_CREDENTIALS
REPOSITORY (Where de events will be sent mongo|api)
MONGO_DB_URI (Mongo URL)
MONGO_DB_USER (Mongo user)
MONGO_DB_PASS (Mongo pass)
MONGO_DB (Mongo db)
MONGO_COLLECTION (Mongo collection)
DRY_RUN (If true don't ACK events from pub/sub, default true)

## Local run ##

go run main.go


## Docker Build ##

```bash
docker build -t pubsub-connect .
```

## Docker Compose ##

Docker compose runs pubsub emulator, creates topic subsciption and publish messages.
Also runs mongo and test api to test repositories.

```bash
docker-compose up
```

## TODO'S ##

- API repository:
    - Custom headers
    - Custom method
    - Timeout
    - Retry
- Mongo repository:
    - json validation
- more repositories
- logger
- filters
- transformers