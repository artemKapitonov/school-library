package migrations

import (
	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/pressly/goose/v3"
)

//go:embed schema/20230603101000__tables.sql
var res embed.FS

// Create runs database migrations.
func Create(db *sqlx.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	goose.SetBaseFS(res)

	if err := goose.Up(db.DB, "schema"); err != nil {
		return err
	}

	return nil
}
