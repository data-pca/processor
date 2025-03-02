package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"processor/config"
	"time"
)

func generateAccessToken(userID string) (*accessTokenData, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims[userKey] = userID
	claims[expKey] = time.Now().Add(config.Cfg.AccessTokenSettings.ExpireMinutes * time.Minute).Unix()

	tokenString, err := token.SignedString([]byte(config.Cfg.AccessTokenSettings.Secret))
	if err != nil {
		return nil, err
	}

	return &accessTokenData{AccessToken: tokenString}, nil
}

func generateRefreshToken(userID string) (*refreshTokenData, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims[userKey] = userID
	claims[expKey] = time.Now().Add(config.Cfg.RefreshTokenSetting.ExpireMinutes * time.Minute).Unix()

	tokenString, err := token.SignedString([]byte(config.Cfg.RefreshTokenSetting.Secret))
	if err != nil {
		return nil, err
	}

	return &refreshTokenData{RefreshToken: tokenString}, nil
}

func extractClaims(tokenString string, secret string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
