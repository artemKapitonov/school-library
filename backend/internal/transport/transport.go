package transport

import "github.com/artemKapitonov/school-library/backend/internal/storage"

type Transport struct {
	BookManager
	StudentManager
	Registrar
}

func New(s *storage.Storage) *Transport {
	return &Transport{
		BookManager:    s.BookStorage,
		StudentManager: s.StudentStorage,
		Registrar: s.RegistrarStorage,

	}
}
