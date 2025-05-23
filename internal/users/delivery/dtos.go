package delivery

import (
	"net/http"
	"reflect"
	"time"

	shareddtos "github.com/gera9/go-blog/internal/shared-models/dtos"
	"github.com/gera9/go-blog/internal/users"
)

type CreatePayload struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Birthdate    time.Time `json:"birthdate"`
}

func (p CreatePayload) Bind(r *http.Request) error {
	return nil
}

func (p CreatePayload) IsValid() bool {
	return !reflect.DeepEqual(p, CreatePayload{})
}

func (p CreatePayload) ToModel() users.User {
	return users.User{
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Email:        p.Email,
		Username:     p.Username,
		PasswordHash: p.PasswordHash,
		Birthdate:    p.Birthdate,
	}
}

type UpdatePayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func (p UpdatePayload) Bind(r *http.Request) error {
	return nil
}

func (p UpdatePayload) IsValid() bool {
	return !reflect.DeepEqual(p, UpdatePayload{})
}

func (p UpdatePayload) ToModel() users.User {
	return users.User{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Username:  p.Username,
		Email:     p.Email,
	}
}

type Response struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Birthdate time.Time `json:"birthdate"`
}

func ToResponse(u users.User) Response {
	return Response{
		Id:        u.Id.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Birthdate: u.Birthdate,
	}
}

func ToListResponse(total int, users []users.User) shareddtos.ListResponse[Response] {
	listResponse := shareddtos.ListResponse[Response]{
		Total: total,
		Items: make([]Response, len(users)),
	}

	for i, u := range users {
		listResponse.Items[i] = ToResponse(u)
	}

	return listResponse
}

// type as workaround for generics bugs with swaggo/swag...
type ListResponse struct {
	Total int        `json:"total"`
	Items []Response `json:"items"`
}
