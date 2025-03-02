package tests

import (
	"github.com/stretchr/testify/assert"
	"processor/config"
	"processor/pkg/jwt"
	"testing"
)

var (
	userID = "12345"
)

func TestGenerateAccessToken(t *testing.T) {
	setConfigValues()

	tokenData, err := jwt.GenerateTokens(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenData.AccessToken)
	assert.NotEmpty(t, tokenData.RefreshToken)
}

func TestValidateAccessToken(t *testing.T) {
	setConfigValues()

	tokens, _ := jwt.GenerateTokens(userID)

	claims, err := jwt.ValidateAccessToken(tokens.AccessToken)

	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
}

func TestUpdateAccessToken(t *testing.T) {
	setConfigValues()

	tokens, _ := jwt.GenerateTokens(userID)

	newAccessToken, err := jwt.UpdateAccessToken(tokens.RefreshToken)

	assert.NoError(t, err)
	assert.NotEmpty(t, newAccessToken.AccessToken)
}

func TestExpiredToken(t *testing.T) {
	setConfigValues()

	config.Cfg.AccessTokenSettings.ExpireMinutes = -1
	config.Cfg.RefreshTokenSetting.ExpireMinutes = -1

	tokenData, err := jwt.GenerateTokens(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenData.AccessToken)
	assert.NotEmpty(t, tokenData.RefreshToken)

	_, err = jwt.ValidateAccessToken(tokenData.AccessToken)
	assert.Error(t, err)
	assert.Equal(t, jwt.ErrTokenExpired, err)
}

func setConfigValues() {
	config.Cfg.AccessTokenSettings.Secret = "accesstestkey"
	config.Cfg.AccessTokenSettings.ExpireMinutes = 10
	config.Cfg.RefreshTokenSetting.Secret = "refreshtestkey"
	config.Cfg.RefreshTokenSetting.ExpireMinutes = 60
}
