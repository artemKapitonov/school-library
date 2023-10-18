package usecase

import "github.com/artemKapitonov/school-library/backend/internal/usecase/storage"

type UseCase struct {
	Class
	Books
	Students
}

func New(s *storage.Storage) *UseCase {
	return &UseCase{}
}
