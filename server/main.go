package main

import (
	"hair-planet/api"
	"log"
	"net/http"

	"encoding/json"

	"github.com/rs/cors"
)

func allName(w http.ResponseWriter, r *http.Request) {
	// Call the GetAllName function to get JSON data
	jsonData, err := api.GetAllName()
	if err != nil {
		// Handle the error (e.g., return an error response)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting data: %v", err)
		return
	}

	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response
	_, err = w.Write(jsonData)
	if err != nil {
		// Handle the error (e.g., log it)
		log.Printf("Error writing response: %v", err)
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	// Decode JSON data from the request body
	var registrationData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsShop   bool   `json:"is_shop"`
	}
	err := json.NewDecoder(r.Body).Decode(&registrationData)
	if err != nil {
		// Handle the decoding error (e.g., return a bad request response)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Call the appropriate registration function
	if registrationData.IsShop {
		err = api.RegisterShop(registrationData.Email, registrationData.Password, registrationData.IsShop)
	} else {
		err = api.RegisterUser(registrationData.Email, registrationData.Password, registrationData.IsShop)
	}

	if err != nil {
		// Handle the registration error (e.g., return an error response)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error registering user/shop: %v", err)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Registration successful"))
}

func handleRequests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/name", allName)
	mux.HandleFunc("/shopDetails", api.ShopDetailsHandler)
	mux.HandleFunc("/createHairdresser", api.CreateHairdresser)
	mux.HandleFunc("/createReservation", api.CreateReservation)
	mux.HandleFunc("/register", registerUser)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":9192", handler))
}

func main() {
	handleRequests()
}
