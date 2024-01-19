package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertIntoHairTest() {
	//RECUPERER EN MODE POST 

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Call insertOneHairTest
	insertOneHairTest(client, "Matthieu", 15)
}

func insertOneHairTest(client *mongo.Client, name string, age int) {
	coll := client.Database("HairPlanet").Collection("HairTest")
	// Add Data
	doc := HairTest{Name: name, Age: age}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

type HairTest struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}
