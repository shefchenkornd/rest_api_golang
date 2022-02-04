package storage

import "github.com/shefchenkornd/rest_api/internal/app/api"

type Storage struct {
	config *api.Config
}

func New(config *api.Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open connection
func (s *Storage) Open() error {
	return nil
}

// Close connection
func (s *Storage) Close() {

}
