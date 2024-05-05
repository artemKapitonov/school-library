package storage

import (
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type Storage struct {
	DB *sqlx.DB
	*BookStorage
	*StudentStorage
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		DB:             db,
		BookStorage:    NewBookStorage(db),
		StudentStorage: NewStudentStorage(db),
	}
}

func Connect() (*sqlx.DB, error) {
	if err := os.MkdirAll("data", 0777); err != nil {
		return nil, err
	}

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
