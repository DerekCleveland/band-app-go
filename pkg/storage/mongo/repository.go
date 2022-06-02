package mongo

import (
	"band-app-go/pkg/input"
	"band-app-go/pkg/insert"
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	dbConn *mongo.Client
}

func ConnectToMongo(e *input.EnvVariables) (*Storage, error) {
	// Set client options
	// clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", e.UsernameStorage, e.PasswordStorage, e.HostStorage))
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s", e.UsernameStorage, e.PasswordStorage, e.HostStorage))

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

	filter := bson.M{
		"bandname":   b.BandName,
		"bandrating": b.BandRating,
		"bandgenre":  b.BandGenre,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Info("Document doesn't exist")
		return nil
	} else if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Document already exists")
	return insert.ErrDuplicate
}
