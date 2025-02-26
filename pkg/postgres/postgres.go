package postgres

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	SuccessfullyInitializedPostgres = "postgres was successfully initialized"
	FailedInitializePostgres        = "postgres failed to initialize"
)

func NewConnector(ctx context.Context, cfg Configurator) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(cfg.GetPgDriverWithConnectionURL())
	if err != nil {
		return nil, err
	}

	if err = conn.PingContext(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}

type Configurator interface {
	GetPgDriverWithConnectionURL() (string, string)
}
