package api

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateHairdresser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var hairdresser Hairdresser
	if err := json.NewDecoder(r.Body).Decode(&hairdresser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()


	// Connect to the MongoDB database
	client, err := connectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	hairdresserCollection := client.Database("HairPlanet").Collection("hairdresser")
	_, err = hairdresserCollection.InsertOne(context.Background(), bson.M{
		"firstName": hairdresser.FirstName,
		"lastName":  hairdresser.LastName,
		"shop_id":   hairdresser.ShopID,
	})
	if err != nil {
		http.Error(w, "Error inserting hairdresser into database", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hairdresser created successfully"))
}
