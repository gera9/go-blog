package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Password  string
	Birthdate time.Time
}

type QueryList struct {
	Limit  int
	Offset int
}
