package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"sync/atomic"

	"github.com/SoraRise/permitted-places/config"
	"github.com/golang-migrate/migrate/database"
)

func init() {
	db := postgres{}
	database.Register("postgres", &db)
	database.Register("postgresql", &db)

}

var (
	miltiStmtDelimiter = []byte(";")

	DefaultMigrationsTable       = "schema_migrations"
	DefaultMultiStatementMaxSize = 10 * 1 << 20
)

var (
	ErrNilConfig = fmt.Errorf("Nil config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
	ErrNoSchema       = fmt.Errorf("no schema")
	ErrDatabaseDirty  = fmt.Errorf("database is dirty")
)

type Postgres struct {
	conn     *sql.Conn
	db       *sql.DB
	isLocked atomic.Bool

	config *config.Config
}

func WithConnection(ctx context.Context, conn *sql.Conn, config *config.Config) (*Postgres, error){
	if config == nil {
		return nil, ErrNilConfig
	}

	if err := conn.PingContext(ctx); err != nil{
		return nil, err
	}

	if config.DatabaseConfig.User
}
