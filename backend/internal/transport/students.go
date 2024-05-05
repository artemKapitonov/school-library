package transport

import (
	"log/slog"

	"github.com/artemKapitonov/school-library/backend/internal/entity"
)

type StudentManager interface {
	GetAllStudents() ([]entity.Student, error)
	SearchStudent(word string) ([]entity.Student, error)
	CreateStudent(student entity.Student) error
	DeleteStudent(id uint64) error
	UpdateStudent(student entity.Student) error
}

func (t *Transport) GetAllStudents() []entity.Student {
	students, err := t.StudentManager.GetAllStudents()
	if err != nil {
		slog.Error("Get All students Error", err)
	}

	return students
}

func (t *Transport) SearchStudent(word string) []entity.Student {
	students, err := t.StudentManager.SearchStudent(word)
	if err != nil {
		slog.Error("Search student Error", err)
	}

	return students
}

func (t *Transport) CreateStudent(s entity.Student) string {
	err := t.StudentManager.CreateStudent(s)
	if err != nil {
		return ""
	}

	return "Success"
}

func (t *Transport) DeleteStudent(id uint64) string {
	err := t.StudentManager.DeleteStudent(id)
	if err != nil {
		return err.Error()
	}

	return "Success"
}

func (t *Transport) UpdateStudent(student entity.Student) string {
	err := t.StudentManager.UpdateStudent(student)
	if err != nil {
		return "ERROR"
	}

	return "Updated"
}
