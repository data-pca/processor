package dto

type GetAuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetAuthResponse struct {
	Token string `json:"token"`
}
