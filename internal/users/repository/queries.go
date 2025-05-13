package repository

const (
	GetById    = "SELECT * FROM users WHERE id = $1"
	Create     = "INSERT INTO users (first_name, last_name, username, email, birthdate, password_hash) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
	List       = "SELECT * FROM users LIMIT $1 OFFSET $2"
	DeleteById = "DELETE FROM users WHERE id = $1"
)
