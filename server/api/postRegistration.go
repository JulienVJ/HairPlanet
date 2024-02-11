package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// RegistrationRequest représente les données d'inscription reçues depuis la requête HTTP
type RegistrationRequest struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	IsShop    bool    `json:"is_shop"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	ShopName  *string `json:"shop_name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
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

	collection := client.Database("HairPlanet").Collection("users")
	userData := bson.M{
		"email":    registrationReq.Email,
		"password": registrationReq.Password,
		"isShop":   registrationReq.IsShop,
	}

	// Ajout des champs optionnels s'ils sont définis
	if registrationReq.FirstName != nil {
		userData["firstName"] = *registrationReq.FirstName
	}
	if registrationReq.LastName != nil {
		userData["lastName"] = *registrationReq.LastName
	}
	if registrationReq.ShopName != nil {
		userData["shopName"] = *registrationReq.ShopName
	}
	if registrationReq.Phone != nil {
		userData["phone"] = *registrationReq.Phone
	}
	if registrationReq.Address != nil {
		userData["address"] = *registrationReq.Address
	}
	_, err = collection.InsertOne(context.Background(), userData)

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
func RegisterUser(email string, password string, isShop bool, firstName *string, lastName *string, shopName *string, phone *string, address *string) error {
    client, err := connectDB()
    if err != nil {
        return err
    }
    defer client.Disconnect(context.Background())

    role := "CLIENT"
    if isShop {
        role = "SHOP"
    }

    // Création de la structure de données pour l'insertion
    userData := bson.M{
        "email":    email,
        "password": password,
        "role":     role,
    }

    // Ajout des champs optionnels s'ils sont définis
    if firstName != nil {
        userData["firstName"] = *firstName
    }
    if lastName != nil {
        userData["lastName"] = *lastName
    }
    if shopName != nil {
        userData["shopName"] = *shopName
    }
    if phone != nil {
        userData["phone"] = *phone
    }
    if address != nil {
        userData["address"] = *address
    }

    // Insertion des données dans la collection "users" de la base de données
    usersCollection := client.Database("HairPlanet").Collection("users")
    _, err = usersCollection.InsertOne(context.Background(), userData)
    if err != nil {
        return err
    }

    return nil
}
