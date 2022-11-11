package client

import (
	"context"
	"fmt"

	"github.com/SoraRise/permitted-places/config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config config.Config) {
	dsn := fmt.Sprintf("")
}
