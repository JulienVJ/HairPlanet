package main

import (
	"encoding/json"
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

func registerUser(w http.ResponseWriter, r *http.Request) {
    // Decode JSON data from the request body
    var registrationData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
        IsShop   bool   `json:"isShop"`
        // Champs facultatifs
        FirstName *string `json:"firstName,omitempty"`
        LastName  *string `json:"lastName,omitempty"`
        ShopName  *string `json:"shopName,omitempty"`
        Phone     *string `json:"phone,omitempty"`
        Address   *string `json:"address,omitempty"`
        Zip       *string `json:"zip,omitempty"`
        City      *string `json:"city,omitempty"`
    }
    err := json.NewDecoder(r.Body).Decode(&registrationData)
    if err != nil {
        // Handle the decoding error (e.g., return a bad request response)
        http.Error(w, "Bad Request", http.StatusBadRequest)
        log.Printf("Error decoding request body: %v", err)
        return
    }

    // Call the appropriate registration function with the password
    err = api.RegisterUser(registrationData.Email, registrationData.Password, registrationData.IsShop, registrationData.FirstName, registrationData.LastName, registrationData.ShopName, registrationData.Phone, registrationData.Address, registrationData.Zip, registrationData.City)

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


func loginUser(w http.ResponseWriter, r *http.Request) {
    // Assurez-vous que la méthode HTTP est POST
    if r.Method != "POST" {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        log.Printf("Invalid HTTP method: %s", r.Method)
        return
    }

    // Decode JSON data from the request body
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    err := json.NewDecoder(r.Body).Decode(&loginData)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        log.Printf("Error decoding login request: %v", err)
        return
    }

    log.Printf("Login request received: %+v", loginData) // Log received login data

    // Appel de la fonction de connexion appropriée pour vérifier les informations d'identification de l'utilisateur
    authToken, err := api.LoginUser(loginData.Email, loginData.Password)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        log.Printf("Error logging in: %v", err)
        return
    }

    // Return the authentication token in the response
    response := map[string]string{"token": authToken}
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        log.Printf("Error encoding response: %v", err)
        return
    }
}

func handleRequests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/name", allName)
	mux.HandleFunc("/shopDetails", api.ShopDetailsHandler)
	mux.HandleFunc("/createHairdresser", api.CreateHairdresser)
	mux.HandleFunc("/createReservation", api.CreateReservation)
	mux.HandleFunc("/login", loginUser)
	mux.HandleFunc("/register", registerUser)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":9192", handler))
}

func main() {
	handleRequests()
}
