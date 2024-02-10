package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// LoginRequest représente les données de connexion reçues depuis la requête HTTP
type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func generateAuthToken(userID primitive.ObjectID) (string, error) {
    // Créer un token avec l'ID utilisateur comme payload
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": userID.Hex(),
    })

    // Signer le token avec une clé secrète (vous devez choisir une clé sécurisée)
    authToken, err := token.SignedString([]byte("votre-clé-secrète"))
    if err != nil {
        return "", err
    }

    return authToken, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Parsez les données JSON de la requête
    var loginReq LoginRequest
    err := json.NewDecoder(r.Body).Decode(&loginReq)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        log.Printf("Error decoding login request: %v", err)
        return
    }

    // Vérifiez les informations d'identification dans la base de données
    client, err := connectDB()
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        log.Printf("Error connecting to database: %v", err)
        return
    }
    defer client.Disconnect(context.Background())

    // Vérifiez si l'utilisateur existe et si le mot de passe correspond
    collection := client.Database("HairPlanet").Collection("users")
    var user bson.M
    err = collection.FindOne(context.Background(), bson.M{"email": loginReq.Email, "password": loginReq.Password}).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        log.Printf("Error finding user: %v", err)
        return
    }

    // Générer un token d'authentification (vous devrez implémenter cette fonctionnalité)
    authToken, err := generateAuthToken(user["_id"].(primitive.ObjectID))
    if err != nil {
        http.Error(w, "Error generating authentication token", http.StatusInternalServerError)
        log.Printf("Error generating authentication token: %v", err)
        return
    }

    // Répondre avec le token d'authentification
    response := map[string]string{"token": authToken}
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

func LoginUser(email string, password string) (string, error) {
    // Connectez-vous à la base de données (vous devrez implémenter cette fonction)
    client, err := connectDB()
    if err != nil {
        return "", err
    }
    defer client.Disconnect(context.Background())

    // Vérifiez les informations d'identification dans la base de données
    collection := client.Database("HairPlanet").Collection("users")
    var user bson.M
    err = collection.FindOne(context.Background(), bson.M{"email": email, "password": password}).Decode(&user)
    if err != nil {
        return "", err
    }

    // Générer un token d'authentification
    authToken, err := generateAuthToken(user["_id"].(primitive.ObjectID))
    if err != nil {
        return "", err
    }

    return authToken, nil
}
