package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sqlx.DB
	*BookStorage
	*StudentStorage
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		BookStorage:    NewBookStorage(db),
		StudentStorage: NewStudentStorage(db),
	}
}

func ConnectToDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "data/lib.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
