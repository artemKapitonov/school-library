package transport

import (
	"fmt"
	"log/slog"

	"github.com/artemKapitonov/school-library/backend/internal/entity"
)

func (t *Transport) GetAllBooks() []entity.Book {
	return []entity.Book{
		entity.Book{
			ID:     10,
			Author: "Pushkin",
			Title:  "Евгений Онегин",
			Year:   10,
			Count:  5,
			Students: []entity.Student{
				entity.Student{
					ID:    100,
					Class: "10B",
					Name:  "Vacay Pupkin",
				},
			},
		},
		entity.Book{
			ID:     10,
			Author: "Тургенев",
			Title:  "Отцы и дети",
			Year:   100,
			Count:  5,
		},
	}
}

func (t *Transport) GetBookByID(word string) entity.Book {
	slog.Info(fmt.Sprintf("Book id is %s", word))
	return entity.Book{
		ID:     10,
		Author: "Pushkin",
		Title:  "Евгений Онегин",
		Year:   10,
		Count:  50,
		Students: []entity.Student{
			entity.Student{
				ID:    100,
				Class: "10B",
				Name:  "Vacay Pupkin",
			},
		},
	}
}

func (t *Transport) CreateBook(book entity.Book) string {
	slog.Info(fmt.Sprintf("Book %s, %s", book.Author, book.Title))
	return "Book was succesfully added"
}

func (t *Transport) DeleteBook(id uint64) string {
	slog.Info(fmt.Sprintf("Book id is %d", id))
	return fmt.Sprintf("Book was succesfully deleted")
}

func (t *Transport) UpdateBook(book entity.Book) string {
	slog.Info(fmt.Sprintf("Book %s, %s", book.Author, book.Title))
	return "Book was succesfully updated"
}
