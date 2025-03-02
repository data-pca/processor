package postgres

const (
	queryGetPassword = `SELECT password FROM auth.accounts WHERE username = $1;`
	queryInsertUser  = `INSERT INTO auth.accounts (created_at, username, password) VALUES ($1, $2, $3) RETURNING id;`
)
