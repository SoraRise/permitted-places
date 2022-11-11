package client

import (
	"context"
	"fmt"
	"internal/poll"

	"log"
	"time"

	"github.com/SoraRise/permitted-places/internal/config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config config.Config) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", &config.DatabaseUserName, &config.UserPassword, &config.DatabaseHost, &config.DatabasePort, &config.DatabaseName)
	err =

	return poll, nil
}
