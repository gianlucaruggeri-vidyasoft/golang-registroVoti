package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongo() *mongo.Database {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	if dbName == "" {
		log.Fatal("ERRORE: La variabile MONGO_DB è vuota")
	}

	clientOptions := options.Client().ApplyURI(uri)
	var client *mongo.Client
	var err error

	for i := 0; i < 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		client, err = mongo.Connect(ctx, clientOptions)

		if err == nil {
			err = client.Ping(ctx, nil)
		}
		cancel()

		if err == nil {
			return client.Database(dbName)
		}

		time.Sleep(3 * time.Second)
	}

	log.Fatalf("Connessione fallita: %v", err)
	return nil
}

