package main

import (
	"encoding/json"
	"hair-planet/api"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func allName(w http.ResponseWriter, r *http.Request) {
	// Appel de la function GetAllName pour récupérer de la data en JSON
	jsonData, err := api.GetAllName()
	if err != nil {
		// Cas d'erreur
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting data: %v", err)
		return
	}

	// Réponse du Header
	w.Header().Set("Content-Type", "application/json")

	// écriture du JSON dans la réponse
	_, err = w.Write(jsonData)
	if err != nil {
		// cas d'erreur
		log.Printf("Error writing response: %v", err)
	}
}

func allShops(w http.ResponseWriter, r *http.Request) {
	// Appel de la function GetShops pour récupérer de la data en JSON
	jsonData, err := api.GetShops()
	if err != nil {
		// cas d'erreur
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting data: %v", err)
		return
	}
	// Réponse du Header
	w.Header().Set("Content-Type", "application/json")
	// écriture du JSON dans la réponse
	_, err = w.Write(jsonData)
	if err != nil {
		// Cas d'erreur
		log.Printf("Error writing response: %v", err)
	}
}

func allReservations(w http.ResponseWriter, r *http.Request) {
	// Appel de la function GetAllReservations pour récupérer de la data en JSON
	jsonData, err := api.GetAllReservations()
	if err != nil {
		// cas d'erreur
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting data: %v", err)
		return
	}
	// Réponse du Header
	w.Header().Set("Content-Type", "application/json")

	// écriture du JSON dans la réponse
	_, err = w.Write(jsonData)
	if err != nil {
		// Cas d'erreur
		log.Printf("Error writing response: %v", err)
	}
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	// Appel de la fonction GetAllUsers pour récupérer de la data en JSON
	jsonData, err := api.GetAllUsers()
	if err != nil {
		// cas d'erreur
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting data: %v", err)
		return
	}

	// Réponse du Header
	w.Header().Set("Content-Type", "application/json")

	// écriture du JSON dans la réponse
	_, err = w.Write(jsonData)
	if err != nil {
		// Cas d'erreur
		log.Printf("Error writing response: %v", err)
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
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
		// Cas d'erreur
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Appeler RegistrationUser
	err = api.RegisterUser(registrationData.Email, registrationData.Password, registrationData.IsShop, registrationData.FirstName, registrationData.LastName, registrationData.ShopName, registrationData.Phone, registrationData.Address, registrationData.Zip, registrationData.City)

	if err != nil {
		// Cas d'erreur
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error registering user/shop: %v", err)
		return
	}

	// Retourne une réponse 200
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

	log.Printf("Login request received: %+v", loginData)

	// Appel de la fonction de connexion appropriée pour vérifier les informations d'identification de l'utilisateur
	authToken, err := api.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		log.Printf("Error logging in: %v", err)
		return
	}

	// Retour avec le token dans la réponse
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
	mux.HandleFunc("/home", allShops)
	mux.HandleFunc("/name", allName)
	mux.HandleFunc("/reservations", allReservations)
	mux.HandleFunc("/users", allUsers)
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
