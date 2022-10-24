package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/MeeL1ght/gopher-rest-api/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	config.ViperConfig()
}

// Get a MongoDB Client
func NewMongoClient() *mongo.Client {
	uri := fmt.Sprintf("%v", viper.Get("db.db_uri"))

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri),
	)

	if err != nil {
		log.Println(err)
	}

	return client
}

// Get a MongoDB Collection
func MongoCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := fmt.Sprintf("%v", viper.Get("db.db_name"))
	collectionName = strings.ToLower(collectionName)
	collectionName = strings.TrimSpace(collectionName)

	return client.Database(dbName).Collection(collectionName)
}
