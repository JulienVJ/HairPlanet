package api

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateReservation(w http.ResponseWriter, r *http.Request) {
	// Vérifie que la reqête est une méthode POST
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

	// Connexion à la BDD
	client, err := connectDB()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	reservationCollection := client.Database("HairPlanet").Collection("reservations")
	_, err = reservationCollection.InsertOne(context.Background(), bson.M{
		"shop_id":     reservation.ShopID,
		"employee_id": reservation.EmployeeID,
		"user_id":     reservation.UserID,
		"date":        reservation.Date,
		"hours":       reservation.Hours,
	})

	if err != nil {
		http.Error(w, "Error inserting reservation into database", http.StatusInternalServerError)
		return
	}

	//Réponse 200
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reservation created successfully"))
}
