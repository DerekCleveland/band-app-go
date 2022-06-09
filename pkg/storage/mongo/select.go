package mongo

import (
	"band-app-go/pkg/listing"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SelectStorageBand takes in band information and selects it from storage
func (s *Storage) SelectStorageBand(b listing.Band) (listing.BandResponse, error) {
	collection := s.dbConn.Database("band-app").Collection("bands")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var tmpBand listing.BandResponse

	err := collection.FindOne(ctx, bson.M{"bandname": b.BandName}).Decode(&tmpBand)
	if err != nil {
		return listing.BandResponse{}, err
	}

	return tmpBand, nil
}
