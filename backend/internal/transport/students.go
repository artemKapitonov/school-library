package transport

import "github.com/artemKapitonov/school-library/backend/internal/entity"

func (t *Transport) GetAllStudents() string {
	return "Student create"
}

func (t *Transport) GetStudentByID(name string) entity.Student {
	return entity.Student{
		ID:    100,
		Class: "10B",
		Name:  "Vacay Pupkin",
		Books: []entity.Book{
			entity.Book{
				ID:     5,
				Author: "Pushkin",
				Title:  "Евгений Онегин",
				Year:   10,
				Count:  5,
			},
			entity.Book{
				ID:     10,
				Author: "Тургенев",
				Title:  "Отцы и дети",
				Year:   100,
				Count:  5,
			},
		},
	}
}

func (t *Transport) CreateStudent(entity.Student) string {
	return "Student create"
}

func (t *Transport) DeleteStudent(id uint64) string {
	return "Student delete"
}

func (t *Transport) UpdateStudent(entity.Student) string {
	return "Student update succesfuly"
}
