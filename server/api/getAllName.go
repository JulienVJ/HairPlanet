package api

import (
	"log"
	"os"
	"fmt"
	"context"
	"encoding/json"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllName() ([]byte, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
		return nil, fmt.Errorf("MONGODB_URI environment variable not set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	coll := client.Database("HairPlanet").Collection("HairTest")

	// Define a filter (empty in this example to get all documents)
	filter := bson.M{}

	// Find documents in the collection
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Error querying MongoDB: %v", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Decode the results
	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Printf("Error decoding MongoDB results: %v", err)
		return nil, err
	}

	// Print or use the results as JSON
	jsonData, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		return nil, err
	}

	return jsonData, nil
}
