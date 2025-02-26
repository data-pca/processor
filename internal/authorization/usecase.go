package authorization

import (
	"context"
	"processor/internal/models/dto"
)

type UseCase interface {
	CheckAuth(ctx context.Context, params dto.GetAuthRequest) (*dto.GetAuthResponse, error)
}
