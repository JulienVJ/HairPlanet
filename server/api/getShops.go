package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents a user document in MongoDB
type User struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func GetShopsHandler(w http.ResponseWriter, r *http.Request) {
	// Set up MongoDB client options
	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Define a slice to store retrieved users
	var shops []User

	// Define the MongoDB collection and filter criteria
	collection := client.Database("HairPlanet").Collection("users")
	log.Println("coll :", collection)

	filter := bson.M{"role": "SHOP"}

	// Query MongoDB for users matching the filter
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)
	log.Println("cur :", cur)

	// Iterate through the results and decode each into a User struct
	for cur.Next(ctx) {
		var user User
		if err := cur.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		shops = append(shops, user)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the retrieved users slice into JSON and write it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shops)
}
