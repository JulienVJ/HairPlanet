package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// RegistrationRequest représente les données d'inscription reçues depuis la requête HTTP
type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsShop   bool   `json:"is_shop"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parsez les données JSON de la requête
	var registrationReq RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&registrationReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding registration request: %v", err)
		return
	}

	// Insérez les données d'inscription dans la base de données
	client, err := connectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Insérez les données dans la collection appropriée en fonction de IsShop
	var collectionName string
	if registrationReq.IsShop {
		collectionName = "shops"
	} else {
		collectionName = "users"
	}

	collection := client.Database("HairPlanet").Collection(collectionName)
	_, err = collection.InsertOne(context.Background(), bson.M{
		"email":    registrationReq.Email,
		"password": registrationReq.Password,
		"isShop":   registrationReq.IsShop,
	})
	if err != nil {
		http.Error(w, "Error inserting registration data", http.StatusInternalServerError)
		log.Printf("Error inserting registration data: %v", err)
		return
	}

	// Répondre avec un message de succès
	response := map[string]string{"message": "Registration successful"}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		log.Printf("Error creating response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// Add debugging statements
func RegisterUser(email string, password string, isShop bool) error {
	fmt.Println("isShop value:", isShop) // Debugging statement

	client, err := connectDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	role := "CLIENT"
	if isShop {
		role = "SHOP"
	}

	usersCollection := client.Database("HairPlanet").Collection("users")
	_, err = usersCollection.InsertOne(context.Background(), bson.M{
		"email":    email,
		"password": password,
		"role":     role,
	})
	if err != nil {
		return err
	}

	return nil
}
