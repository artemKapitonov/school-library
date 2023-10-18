package storage

import "github.com/jackc/pgx/v5/pgxpool"

type ClassPsql struct {
	db *pgxpool.Pool
}

func NewClassPsql(db *pgxpool.Pool) *ClassPsql {
	return &ClassPsql{db: db}
}
