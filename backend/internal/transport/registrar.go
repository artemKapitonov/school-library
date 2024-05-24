package transport

import (
	"database/sql"
	"log/slog"

	"github.com/artemKapitonov/school-library/backend/internal/transport/messages"
)

type Registrar interface {
	RegisterStudentBook(bookID, stdentID uint64) error
	UnregisterStudentBook(bookID, studentID uint64) error
}

func (t *Transport) RegisterStudentBook(bookID, studentID uint64) string {
	err := t.Registrar.RegisterStudentBook(bookID, studentID)
	if err == sql.ErrNoRows {
		return messages.StudentOrBookNotFound
	}

	if err != nil {
		slog.Error("Can't register Error:", err)
		return messages.ErrRegister
	}

	return messages.BookRegistered
}

func (t *Transport) UnregisterStudentBook(bookID, studentID uint64) string {
	err := t.Registrar.UnregisterStudentBook(bookID, studentID)
	if err == sql.ErrNoRows {
		return messages.StudentOrBookNotFound
	}

	if err != nil {
		slog.Error("Can't register Error:", err)
		return messages.ErrUnregister
	}

	return messages.BookUnregistered
}
