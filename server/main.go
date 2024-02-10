package main

import (
	"hair-planet/api"
	"log"
	"net/http"

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

func handleRequests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/name", allName)
	mux.HandleFunc("/shopDetails", api.ShopDetailsHandler)
	mux.HandleFunc("/createHairdresser", api.CreateHairdresser)
	mux.HandleFunc("/createReservation", api.CreateReservation)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":9192", handler))
}

func main() {
	handleRequests()
}
