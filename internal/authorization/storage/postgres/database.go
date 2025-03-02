package postgres

import (
	"context"
	"github.com/dromara/carbon/v2"
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

func (d database) SignIn(ctx context.Context, params dao.SignInRequest) (*dao.SignInResponse, error) {
	var response dao.SignInResponse

	if err := d.db.GetContext(ctx, queryGetPassword, params.Username); err != nil {
		return nil, err
	}

	return &response, nil
}

func (d database) SignUp(ctx context.Context, params dao.SignUpRequest) (*dao.SignUpResponse, error) {
	var response dao.SignUpResponse

	if err := d.db.GetContext(ctx, &response, queryInsertUser, carbon.Now(), params.Username, params.Password); err != nil {
		return nil, err
	}

	return &response, nil
}
