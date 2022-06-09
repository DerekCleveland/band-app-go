package listing

import "errors"

type Band struct {
	BandName string
}

type BandResponse struct {
	BandName   string  `json:"name"`
	BandRating float64 `json:"rating"`
	BandGenre  string  `json:"genre"`
}

// ErrNotFound is used when a band is not found.
var ErrNotFound = errors.New("band not found")

// Repository provides access to the database storage (the functions in the storage package)
type Repository interface {
	// SelectStorageBand inserts band information into storage
	SelectStorageBand(Band) (BandResponse, error)
}

// Service provides endpoint with adding operations
type Service interface {
	SelectBand(Band) (BandResponse, error)
}

type service struct {
	r Repository
}

// NewService creates a listing of service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// InsertBand inserts a band into storage
func (s *service) SelectBand(b Band) (BandResponse, error) {
	return s.r.SelectStorageBand(b)
}
