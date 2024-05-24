package storage

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RegistrarStorage struct {
	db *sqlx.DB
}

func NewRegistrarStorage(db *sqlx.DB) *RegistrarStorage {
	return &RegistrarStorage{db: db}
}

func (s *RegistrarStorage) RegisterStudentBook(bookID, studentID uint64) error {
	var query = fmt.Sprintf("INSERT INTO students_books (book_id, student_id) VALUES (%d, %d);", bookID, studentID)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *RegistrarStorage) UnregisterStudentBook(bookID, studentID uint64) error {
	var query = fmt.Sprintf("DELETE FROM students_books WHERE book_id = %d AND student_id = %d", bookID, studentID)

	result, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Book not deleted")
	}

	return nil
}
