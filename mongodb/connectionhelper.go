package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	username   = ""
	database   = ""
	password   = ""
	db         = "PROD"
	collection = "exchange-rates"
)

var connectionString = fmt.Sprintf(
	"mongodb://%s:%s@%s/?ssl=true&retrywrites=false",
	username,
	password,
	database)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(connectionString)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)

		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
