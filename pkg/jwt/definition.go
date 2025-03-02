package jwt

import "errors"

const (
	userKey = "user_id"
	expKey  = "exp"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrMissingUserID = errors.New("missing user_id")
	ErrTokenExpired  = errors.New("token is expired")
)

// ClaimsPayload contains all embaded system data
type ClaimsPayload struct {
	UserID string
}

// AccessToken - keeps token used for passing authorization and parsing client info
type accessTokenData struct {
	AccessToken string `json:"access_token"`
}

// RefreshToken - keeps token used for updating AccessToken after its expire
type refreshTokenData struct {
	RefreshToken string `json:"refresh_token"`
}

// Tokens - used for grouping AccessToken and RefreshToken for initial client authorization response
type Tokens struct {
	*accessTokenData
	*refreshTokenData
}
