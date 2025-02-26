package usecase

import (
	"context"
	"processor/internal/authorization"
	"processor/internal/models/dao"
	"processor/internal/models/dto"
)

type usecase struct {
	storage authorization.Storage
}

func New(storage authorization.Storage) authorization.UseCase {
	return &usecase{storage: storage}
}

func (u usecase) CheckAuth(ctx context.Context, params dto.GetAuthRequest) (*dto.GetAuthResponse, error) {
	response, err := u.storage.CheckAuth(ctx, dao.GetAuthRequest(params))
	if err != nil {
		return nil, err
	}

	return &dto.GetAuthResponse{Token: response.Token}, nil
}
