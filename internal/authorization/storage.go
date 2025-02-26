package authorization

import (
	"context"
	"processor/internal/models/dao"
)

type Storage interface {
	CheckAuth(ctx context.Context, params dao.GetAuthRequest) (*dao.GetAuthResponse, error)
}
