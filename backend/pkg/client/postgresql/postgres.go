package postgresql

import (
	"context"
	"fmt"
	"github.com/artemKapitonov/school-library/backend/pkg/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Client is an interface for PostgresSQL client operations
type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Ping(ctx context.Context) error
	Close()
}

// Config represents the configuration for connecting to PostgresSQL
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// ConnectionConfig creates a connection configuration for PostgresSQL
func connectionConfig(cfg Config) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ConnectToDB connects to the PostgresSQL database
func ConnectToDB(ctx context.Context, cfg Config) (db *pgxpool.Pool, err error) {
	var maxAttempts = 5
	connCfg, err := connectionConfig(cfg)
	if err != nil {
		return nil, err
	}

	err = utils.DoWithTries(
		func() error {
			ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()

			db, err = pgxpool.NewWithConfig(ctx, connCfg)
			if err != nil {
				return err
			}

			if err := db.Ping(ctx); err != nil {
				return err
			}

			return nil
		}, maxAttempts, 5*time.Second)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
