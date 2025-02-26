package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"processor/internal/authorization"
	"processor/internal/models/dao"
)

type database struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) authorization.Storage {
	return &database{db: db}
}

func (d database) CheckAuth(ctx context.Context, params dao.GetAuthRequest) (*dao.GetAuthResponse, error) {
	var response dao.GetAuthResponse

	query := "SELECT * FROM auth WHERE login = $1 AND password = $2"

	if err := d.db.GetContext(ctx, &response, query, params.Login, params.Password); err != nil {
		return nil, err
	}

	return &response, nil
}
