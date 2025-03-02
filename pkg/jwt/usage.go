package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"processor/config"
)

func GenerateTokens(userID string) (*Tokens, error) {
	accessToken, err := generateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		accessToken,
		refreshToken,
	}, nil
}

func ValidateAccessToken(accessToken string) (ClaimsPayload, error) {
	var payload ClaimsPayload

	claims, err := extractClaims(accessToken, config.Cfg.AccessTokenSettings.Secret)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return payload, ErrTokenExpired
		}
		return payload, err
	}

	claimsUserID, ok := claims[userKey]
	if !ok {
		return payload, ErrMissingUserID
	}

	userID, ok := claimsUserID.(string)
	if !ok {
		return payload, ErrMissingUserID
	}

	payload.UserID = userID

	return payload, nil
}

func UpdateAccessToken(refreshToken string) (*accessTokenData, error) {
	claims, err := extractClaims(refreshToken, config.Cfg.RefreshTokenSetting.Secret)
	if err != nil {
		return nil, err
	}

	claimsUserID, ok := claims[userKey]
	if !ok {
		return nil, ErrMissingUserID
	}

	userID, ok := claimsUserID.(string)
	if !ok {
		return nil, ErrMissingUserID
	}

	return generateAccessToken(userID)
}
