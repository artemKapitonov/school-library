package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Storage struct {
	*ClassPsql
	*BooksPsql
	*StudentsPsql
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		ClassPsql:    NewClassPsql(db),
		BooksPsql:    NewBooksPsql(db),
		StudentsPsql: NewStudentsPsql(db),
	}
}
