package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/artemKapitonov/school-library/backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BookStorage struct {
	db *sqlx.DB
}

func NewBookStorage(db *sqlx.DB) *BookStorage {
	return &BookStorage{db: db}
}

func (s *BookStorage) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	var query = "SELECT id, author, title, year, count FROM books;"

	err := s.db.Select(&books, query)
	if err != nil {
		return nil, err
	}

	s.setStudentsForBooks(books)

	return books, nil
}

func (s *BookStorage) SearchBook(word string) ([]entity.Book, error) {
	var books []entity.Book

	word = "%" + word + "%"

	var query = fmt.Sprintf(`SELECT id, author, title, year, count FROM books
	WHERE id LIKE '%s' OR title LIKE '%s' OR author LIKE '%s';`, word, word, word)

	err := s.db.Select(&books, query)
	if err != nil {
		return nil, err
	}

	s.setStudentsForBooks(books)

	return books, nil
}

func (s *BookStorage) CreateBook(b entity.Book) error {
	query := fmt.Sprintf(
		"INSERT INTO books (author, title, year, count) VALUES ('%s', '%s', %d, %d);",
		b.Author, b.Title, b.Year, b.Count,
	)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookStorage) DeleteBook(id uint64) error {
	var query = fmt.Sprintf("DELETE FROM books WHERE id = %d", id)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookStorage) UpdateBook(book entity.Book) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	argsID := 1

	if book.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsID))
		args = append(args, book.Title)
		argsID++
	}

	if book.Author != "" {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argsID))
		args = append(args, book.Author)
		argsID++
	}

	if book.Year != 0 {
		setValues = append(setValues, fmt.Sprintf("year=$%d", argsID))
		args = append(args, book.Year)
		argsID++
	}

	if book.Count != 0 {
		setValues = append(setValues, fmt.Sprintf("count=$%d", argsID))
		args = append(args, book.Count)
		argsID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE books SET %s WHERE id=$%d;", setQuery, argsID,
	)

	slog.Info(query)

	args = append(args, book.ID)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}


func (s *BookStorage) getStudentsByBookID(bookID uint64) ([]entity.Student, error) {
	var students []entity.Student

	getStudentsQuery := fmt.Sprintf(
		`SELECT s.id, s.class, s.name FROM students s 
		INNER JOIN students_books sb on sb.student_id = s.id WHERE sb.book_id = %d;`,
		bookID,
	)

	err := s.db.Select(&students, getStudentsQuery)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return students, nil
}

func (s *BookStorage) setStudentsForBooks(books []entity.Book) error {
	var err error

	for i := range books {
		books[i].Students, err = s.getStudentsByBookID(books[i].ID)
		if err != nil {
			return err
		}
	}

	return nil
}
