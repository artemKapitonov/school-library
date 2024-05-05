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

type StudentStorage struct {
	db *sqlx.DB
}

func NewStudentStorage(db *sqlx.DB) *StudentStorage {
	return &StudentStorage{db: db}
}

func (s *StudentStorage) GetAllStudents() ([]entity.Student, error) {
	var students []entity.Student
	var query = "SELECT id, class, name FROM students"

	err := s.db.Select(&students, query)
	if err != nil {
		return nil, err
	}

	s.setBooksByStudents(students)

	return students, nil
}

func (s *StudentStorage) SearchStudent(word string) ([]entity.Student, error) {
	var students []entity.Student

	word = "%" + word + "%"

	var query = fmt.Sprintf(`SELECT id, class, name FROM students
	WHERE id LIKE '%s' OR class LIKE '%s' OR name LIKE '%s';`, word, word, word)

	err := s.db.Select(&students, query)
	if err != nil {
		return nil, err
	}

	err = s.setBooksByStudents(students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *StudentStorage) CreateStudent(student entity.Student) error {
	var query = fmt.Sprintf("INSERT INTO students (class, name) VALUES ('%s', '%s');", student.Class, student.Name)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentStorage) DeleteStudent(id uint64) error {
	var query = fmt.Sprintf("DELETE FROM students WHERE id = %d;", id)

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentStorage) UpdateStudent(student entity.Student) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	argsID := 1

	if student.Class != "" {
		setValues = append(setValues, fmt.Sprintf("class=$%d", argsID))
		args = append(args, student.Class)
		argsID++
	}

	if student.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argsID))
		args = append(args, student.Name)
		argsID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE students SET %s WHERE id=$%d;", setQuery, argsID,
	)

	slog.Info(query)

	args = append(args, student.ID)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentStorage) getBooksByStudentID(id uint64) ([]entity.Book, error) {
	var books []entity.Book

	getStudentsQuery := fmt.Sprintf(
		`SELECT b.id, b.author, b.title, b.year, b.count FROM books b 
		INNER JOIN students_books sb on sb.book_id = b.id WHERE sb.student_id = %d;`,
		id,
	)

	err := s.db.Select(&books, getStudentsQuery)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return books, nil
}

func (s *StudentStorage) setBooksByStudents(students []entity.Student) error {
	var err error

	for i := range students {
		students[i].Books, err = s.getBooksByStudentID(students[i].ID)
		if err != nil {
			return err
		}
	}

	return nil
}
