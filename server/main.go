package main

import (
	"encoding/json"
	"fmt"
	"hair-planet/api"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hompeage Hello")
}

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
		IsShop   bool   `json:"isShop"`
	}
	err := json.NewDecoder(r.Body).Decode(&registrationData)
	if err != nil {
		// Handle the decoding error (e.g., return a bad request response)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}
	// Call the appropriate registration function
	err = api.RegisterUser(registrationData.Email, registrationData.Password, registrationData.IsShop)

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
	http.HandleFunc("/", homepage)
	mux := http.NewServeMux()
	mux.HandleFunc("/name", allName)
	mux.HandleFunc("/register", registerUser)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":9192", handler))
}

func main() {
	handleRequests()
}
