package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SoraRise/permitted-places/internal/config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config config.Config) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.DatabaseUserName, config.UserPassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	err = DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, &config)
	if err != nil {
		log.Fatal(" error do with tries postgresql ")
	}

	return pool, err
}

func DoWithTries(fn func() error, config *config.Config) (err error) {
	for config.MaxAttemptsConnections > 0 {
		if err = fn(); err != nil {
			time.Sleep(5 * time.Second)
			config.MaxAttemptsConnections--
			continue
		}
		return nil
	}
	return
}
