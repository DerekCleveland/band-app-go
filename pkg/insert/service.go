package insert

import "errors"

type Band struct {
	BandName   string
	BandRating float64
	BandGenre  string
	// BandSongCount  int
	// BandAlbumCount int
}

// ErrDuplicate is used when a band already exists.
var ErrDuplicate = errors.New("band already exists")

// Repository provides access to the database storage (the functions in the storage package)
type Repository interface {
	// InsertStorageBand inserts band information into storage
	InsertStorageBand(Band) error

	// CheckStorageIfBandExists queries storage to see if a band has already been inserted. Returns an error if one has
	CheckStorageIfBandExists(Band) error
}

// Service provides endpoint with adding operations
type Service interface {
	InsertBand(Band) error
	CheckIfBandExists(Band) error
}

type service struct {
	r Repository
}

// NewService creates a listing of service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// InsertBand inserts a band into storage
func (s *service) InsertBand(b Band) error {
	return s.r.InsertStorageBand(b)
}

func (s *service) CheckIfBandExists(b Band) error {
	return s.r.CheckStorageIfBandExists(b)
}
