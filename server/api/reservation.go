package api

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type Reservation struct {
	ID         string `json:"_id" bson:"_id"`
	ShopID     string `json:"shop_id" bson:"shop_id"`
	EmployeeID string `json:"employee_id" bson:"employee_id"`
	UserID     string `json:"user_id" bson:"user_id"`
	Date       string `json:"date" bson:"date"`
	Hours      string `json:"hours" bson:"hours"`
}


func CreateReservation(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reservation Reservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
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

	reservationCollection := client.Database("HairPlanet").Collection("reservations")
	_, err = reservationCollection.InsertOne(context.Background(), bson.M{
		"shop_id": reservation.ShopID,
		"employee_id":reservation.EmployeeID,
		"user_id":reservation.UserID,
		"date":reservation.Date,
		"hours":reservation.Hours,
	})

	if err != nil {
		http.Error(w, "Error inserting reservation into database", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reservation created successfully"))
}