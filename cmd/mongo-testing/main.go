package main

import (
	"band-app-go/pkg/util"
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
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
	// e, err := input.GetEnvVariables()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	log.Info().Msg("✅ Loaded config - Got enviroment variables")

	log.Info().Msg("user: " + config.MongoUsername)
	log.Info().Msg("password: " + config.MongoPassword)
	log.Info().Msg("host: " + config.MongoHost)
	log.Info().Msg("port: " + config.MongoPort)

	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s@%s:%s", config.MongoScheme, config.MongoUsername, config.MongoPassword, config.MongoHost, config.MongoPort))
	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", e.UsernameStorage, e.PasswordStorage, e.HostStorage))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to MongoDB")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ping MongoDB")
	}

	log.Info().Msg("✅ Connected to MongoDB")

	collection := client.Database("band-app").Collection("test-bands")

	// INSERTING DOC
	bandDocument := InsertBand{
		BandName:   "Thornhill",
		BandRating: 9.7,
		BandGenre:  "Metal",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insertResult, err := collection.InsertOne(ctx, bandDocument)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to insert document")
	}

	log.Info().Msgf("Inserted a single document: %+v", insertResult.InsertedID)

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
		log.Info().Msg("Document doesn't exist")
		bandInStorage = false
	} else if err != nil {
		log.Fatal().Err(err).Msg("Failed to find document")
	}

	log.Info().Msg("✅ Found document")
	log.Info().Msg("Bandname: " + result.BandName)
	log.Info().Msg("Bandgenre: " + result.BandGenre)
	log.Info().Msg(fmt.Sprintf("Bandrating: %f", result.BandRating))
	log.Info().Msg(fmt.Sprintf("BandInStorage: %t", bandInStorage))
}
