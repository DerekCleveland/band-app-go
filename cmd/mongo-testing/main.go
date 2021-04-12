package main

import (
	"band-app-go/pkg/input"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertBand struct {
	BandName   string
	BandRating float64
	BandGenre  string
}

func main() {
	e, err := input.GetEnvVariables()
	if err != nil {
		log.Fatal(err)
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", e.UsernameStorage, e.PasswordStorage, e.HostStorage))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("band-app").Collection("test-bands")

	// INSERTING DOC
	// bandDocument := InsertBand{
	// 	BandName:   "Thornhill",
	// 	BandRating: 9.7,
	// 	BandGenre:  "Metal",
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// insertResult, err := collection.InsertOne(ctx, bandDocument)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("inserted single document:", insertResult.InsertedID)

	// FIND SINGLE DOC
	var result InsertBand
	var bandInStorage = true

	filter := bson.M{
		"bandname":   "Thornhill",
		"bandrating": 9.7,
		"bandgenre":  "Metal",
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Println(err)
		bandInStorage = false
	} else if err != nil {
		log.Println(err)
	}

	fmt.Println("bandname:", result.BandName)
	fmt.Println("bandgenre:", result.BandGenre)
	fmt.Println("bandrating:", result.BandRating)
	fmt.Println("bandInStorage:", bandInStorage)
}
