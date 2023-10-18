package storage

import "github.com/jackc/pgx/v5/pgxpool"

type BooksPsql struct {
	db *pgxpool.Pool
}

func NewBooksPsql(db *pgxpool.Pool) *BooksPsql {
	return &BooksPsql{db: db}
}
