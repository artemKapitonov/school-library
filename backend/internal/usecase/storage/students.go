package storage

import "github.com/jackc/pgx/v5/pgxpool"

type StudentsPsql struct {
	db *pgxpool.Pool
}

func NewStudentsPsql(db *pgxpool.Pool) *StudentsPsql {
	return &StudentsPsql{db: db}
}
