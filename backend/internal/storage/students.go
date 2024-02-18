package storage

import "github.com/jmoiron/sqlx"

type StudentStorage struct {
	db *sqlx.DB
}

func NewStudentStorage(db *sqlx.DB) *StudentStorage {
	return &StudentStorage{db: db}
}

func (s *StudentStorage) GetAllStudents() {
	panic("implement me")
}

func (s *StudentStorage) GetStudentByID() {
	panic("implement me")
}

func (s *StudentStorage) CreateStudent() {
	panic("implement me")
}

func (s *StudentStorage) DeleteStudent() {
	panic("implement me")
}

func (s *StudentStorage) UpdateStudent() {
	panic("implement me")
}
