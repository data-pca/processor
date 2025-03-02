package authorization

import (
	"context"
	"processor/internal/models/dto"
	"processor/pkg/jwt"
)

type UseCase interface {
	SignIn(ctx context.Context, params dto.SignInRequest) (*jwt.Tokens, error)
	SignUp(ctx context.Context, params dto.SignUpRequest) (*jwt.Tokens, error)
}
