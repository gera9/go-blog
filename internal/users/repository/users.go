package repository

import (
	"time"

	"github.com/gera9/go-blog/internal/users"
	"github.com/google/uuid"
)

var inMemoryUsers = []users.User{
	{
		Id:        uuid.New(),
		FirstName: "Alice",
		LastName:  "Smith",
		Username:  "alice123",
		Password:  "passAlice!",
		Birthdate: time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:        uuid.New(),
		FirstName: "Bob",
		LastName:  "Johnson",
		Username:  "bobbyj",
		Password:  "bobPass2024",
		Birthdate: time.Date(1985, 8, 20, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:        uuid.New(),
		FirstName: "Carol",
		LastName:  "Williams",
		Username:  "carol_w",
		Password:  "carolSecure",
		Birthdate: time.Date(1992, 12, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:        uuid.New(),
		FirstName: "David",
		LastName:  "Brown",
		Username:  "david.b",
		Password:  "david1234",
		Birthdate: time.Date(1998, 3, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:        uuid.New(),
		FirstName: "Eve",
		LastName:  "Davis",
		Username:  "eveD",
		Password:  "ev3Secure!",
		Birthdate: time.Date(2000, 7, 30, 0, 0, 0, 0, time.UTC),
	},
}
