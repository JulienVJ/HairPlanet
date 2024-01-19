package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListByName() ([]byte, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return nil, err
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
		return nil, fmt.Errorf("MONGODB_URI environment variable not set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	coll := client.Database("HairPlanet").Collection("HairTest")
	name := "Martin"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the name %s\n", name)
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
