package storage

import "github.com/jmoiron/sqlx"

type BookStorage struct {
	db *sqlx.DB
}

func NewBookStorage(db *sqlx.DB) *BookStorage {
	return &BookStorage{db: db}
}

func (s *BookStorage) GetAllBooks() {
	panic("implement me")
}

func (s *BookStorage) GetBookByID() {
	panic("implement me")
}

func (s *BookStorage) CreateBook() {
	panic("implement me")
}

func (s *BookStorage) DeleteBook() {
	panic("implement me")
}

func (s *BookStorage) UpdateBook() {
	panic("implement me")
}
