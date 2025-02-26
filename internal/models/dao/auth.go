package dao

type GetAuthRequest struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type GetAuthResponse struct {
	Token string `db:"token"`
}
