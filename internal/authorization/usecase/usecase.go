package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"processor/internal/authorization"
	"processor/internal/models/dao"
	"processor/internal/models/dto"
	"processor/pkg/jwt"
	"strconv"
)

type usecase struct {
	storage authorization.Storage
}

func New(storage authorization.Storage) authorization.UseCase {
	return &usecase{storage: storage}
}

func (u usecase) SignIn(ctx context.Context, params dto.SignInRequest) (*jwt.Tokens, error) {
	response, err := u.storage.SignIn(ctx, dao.SignInRequest(params))
	if err != nil {
		return nil, err
	}

	if err = comparePasswordHash(params.Password, response.Password); err != nil {
		return nil, authorization.ErrPasswordsNotMatch
	}

	return jwt.GenerateTokens(strconv.Itoa(response.ID))
}

func (u usecase) SignUp(ctx context.Context, params dto.SignUpRequest) (*jwt.Tokens, error) {
	hash, err := generatePasswordHash(params.Password)
	if err != nil {
		return nil, err
	}

	request := dao.SignUpRequest{
		Username: params.Username,
		Password: hash,
	}

	response, err := u.storage.SignUp(ctx, request)
	if err != nil {
		return nil, err
	}

	return jwt.GenerateTokens(strconv.Itoa(response.UserID))
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
