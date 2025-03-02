package authorization

import (
	"context"
	"processor/internal/models/dao"
)

type Storage interface {
	SignIn(ctx context.Context, params dao.SignInRequest) (*dao.SignInResponse, error)
	SignUp(ctx context.Context, params dao.SignUpRequest) (*dao.SignUpResponse, error)
}
