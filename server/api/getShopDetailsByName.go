package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Shop struct represents a user document in MongoDB
type Shop struct {
	ID   string `bson:"_id"`
	name string `bson:"name"`
	phone string `bson:"phone"`
	address string `bson:"address"`
	zip string `bson:"zip"`
	city string `bson:"city"`
}

// Hairdresser struct represents a hairdresser document in MongoDB
type Hairdresser struct {
	ID      string `bson:"_id"`
	ShopID  string `bson:"shop_id"`
	FirstName    string `bson:"firstName"`
	LastName string `bson:"lastName"`
}

func connectDB() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
    }
	uri := os.Getenv("MONGODB_URI")
	log.Println("User :", uri)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func AllDetailsShopByName(shopName string) ([]byte, error) {
	client, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	usersCollection := client.Database("HairPlanet").Collection("users")
	var user bson.M
	err = usersCollection.FindOne(context.Background(), bson.D{{"name", shopName}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	log.Println("User :", user)

	userID := user["_id"].(primitive.ObjectID)
	shopID := userID.Hex()

	log.Println("User ID:", userID)
	log.Println("Shop ID:", shopID)


	hairdresserCollection := client.Database("HairPlanet").Collection("hairdresser")
	cursor, err := hairdresserCollection.Find(context.TODO(), bson.D{{"shop_id", shopID}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var hairdressers []Hairdresser
	if err := cursor.All(context.Background(), &hairdressers); err != nil {
		return nil, err
	}

	combinedData := map[string]interface{}{
		"user":         user,
		"hairdressers": hairdressers,
	}

	jsonData, err := json.MarshalIndent(combinedData, "", "    ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func ShopDetailsHandler(w http.ResponseWriter, r *http.Request) {
	shopName := r.URL.Query().Get("name")

	jsonData, err := AllDetailsShopByName(shopName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting data: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
