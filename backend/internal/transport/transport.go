package transport

import "github.com/artemKapitonov/school-library/backend/internal/storage"

type Transport struct {
	s *storage.Storage
}

func New(s *storage.Storage) *Transport {
	return &Transport{s: s}
}
