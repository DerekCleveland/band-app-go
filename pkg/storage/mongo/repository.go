package mongo

import (
	"band-app-go/pkg/insert"
	"band-app-go/pkg/util"
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	dbConn *mongo.Client
}

func ConnectToMongo(c *util.Config) (*Storage, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s@%s:%s", c.MongoScheme, c.MongoUsername, c.MongoPassword, c.MongoHost, c.MongoPort))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return &Storage{dbConn: client}, nil
}

// InsertStorageBand takes in band information and inserts it into storage
func (s *Storage) InsertStorageBand(b insert.Band) error {
	collection := s.dbConn.Database("band-app").Collection("bands")

	document := InsertBand{
		BandName:   b.BandName,
		BandRating: b.BandRating,
		BandGenre:  b.BandGenre,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CheckStorageIfBandExists(b insert.Band) error {
	collection := s.dbConn.Database("band-app").Collection("bands")

	var result InsertBand

	// TODO Should only filter off of bandname. Other two options should be handled by updates
	filter := bson.M{
		"bandname":   b.BandName,
		"bandrating": b.BandRating,
		"bandgenre":  b.BandGenre,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Debug().Msg("Document doesn't exist")
		return nil
	} else if err != nil {
		log.Error().Err(err).Msg("Failed to find document")
		return err
	}

	log.Info().Msgf("Document already exists: %+v", result)
	return insert.ErrDuplicate
}
