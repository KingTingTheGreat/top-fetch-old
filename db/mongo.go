package db

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBUser struct {
	SpotifyId    string `bson:"spotifyId"`
	Id           string `bson:"id"`
	AccessToken  string
	RefreshToken string
}

func generateId() string {
	randStr := func() string {
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		b := make([]byte, 32)
		for i := range b {
			b[i] = chars[rand.Intn(len(chars))]
		}
		return string(b)
	}

	for {
		id := randStr()
		_, err := GetUserById(id)
		if err != nil {
			return id
		}
	}
}

func ConnectDB() *mongo.Client {
	godotenv.Load()
	bg := context.Background()
	wT, cancel := context.WithTimeout(bg, 10000*time.Millisecond)
	defer func() { cancel() }()
	client, err := mongo.Connect(wT, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(wT, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to mongodb")
	return client
}

var db *mongo.Client = ConnectDB()
var userColletion *mongo.Collection = getCollection(os.Getenv("COLLECTION_NAME"))

func getCollection(collectionName string) *mongo.Collection {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	return db.Database(os.Getenv("DB_NAME") + env).Collection(collectionName)
}

func GetUserById(id string) (DBUser, error) {
	var user DBUser
	err := userColletion.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	return user, err
}

func GetUserBySpotifyId(spotifyId string) (DBUser, error) {
	var user DBUser
	err := userColletion.FindOne(context.Background(), bson.M{"spotifyId": spotifyId}).Decode(&user)
	return user, err
}

func InsertUser(user DBUser) (string, error) {
	user.Id = generateId()

	_, err := userColletion.InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return user.Id, nil
}
