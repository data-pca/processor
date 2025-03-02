package dao

type (
	SignInRequest struct {
		Username string
		Password string
	}

	SignInResponse struct {
		ID       int    `db:"id"`
		Password string `db:"password"`
	}
)

//

type (
	SignUpRequest struct {
		Username string
		Password string
	}

	SignUpResponse struct {
		UserID int `db:"id"`
	}
)
