package transport

import (
	"database/sql"
	"log/slog"

	"github.com/artemKapitonov/school-library/backend/internal/entity"
	"github.com/artemKapitonov/school-library/backend/internal/transport/messages"
)

type BookManager interface {
	GetAllBooks() ([]entity.Book, error)
	SearchBook(word string) ([]entity.Book, error)
	CreateBook(book entity.Book) error
	DeleteBook(id uint64) error
	UpdateBook(entity.Book) error
}



func (t *Transport) GetAllBooks() []entity.Book {
	books, err := t.BookManager.GetAllBooks()
	if err != nil {
		slog.Error("Get all book error", err)
	}

	return books
}

func (t *Transport) SearchBook(word string) []entity.Book {
	books, err := t.BookManager.SearchBook(word)
	if err != nil {
		slog.Error("Search book Error", err)
	}

	return books
}

func (t *Transport) CreateBook(book entity.Book) string {
	err := t.BookManager.CreateBook(book)
	if err != nil {
		slog.Error("Can't create book", err)
		return messages.ErrBookNotCreated
	}

	return messages.BookCreated
}

func (t *Transport) DeleteBook(id uint64) string {
	err := t.BookManager.DeleteBook(id)
	if err == sql.ErrNoRows {
		return messages.BookNotFound
	}

	if err != nil {
		slog.Error("Can't delete book", err)
		return messages.ErrBookNotDeleted
	}

	return messages.BookDeleted
}

func (t *Transport) UpdateBook(book entity.Book) string {
	err := t.BookManager.UpdateBook(book)
	if err == sql.ErrNoRows {
		return messages.BookNotFound
	}

	if err != nil {
		slog.Error("Can't update book", err)
		return messages.ErrBookNotUpdated
	}

	return messages.BookUpdated
}


