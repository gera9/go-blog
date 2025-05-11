package users

import (
	"net/http"
	"reflect"
	"time"

	shareddtos "github.com/gera9/go-blog/internal/shared-models/dtos"
)

type CreatePayload struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Birthdate time.Time `json:"birthdate"`
}

func (p CreatePayload) Bind(r *http.Request) error {
	return nil
}

func (p CreatePayload) IsValid() bool {
	return !reflect.DeepEqual(p, CreatePayload{})
}

func (p CreatePayload) ToModel() User {
	return User{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Username:  p.Username,
		Password:  p.Password,
		Birthdate: p.Birthdate,
	}
}

type UpdatePayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func (p UpdatePayload) Bind(r *http.Request) error {
	return nil
}

func (p UpdatePayload) IsValid() bool {
	return !reflect.DeepEqual(p, UpdatePayload{})
}

func (p UpdatePayload) ToModel() User {
	return User{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Username:  p.Username,
	}
}

type Response struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Birthdate time.Time `json:"birthdate"`
}

func ToResponse(u User) Response {
	return Response{
		Id:        u.Id.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Birthdate: u.Birthdate,
	}
}

func ToListResponse(total int, users []User) shareddtos.ListResponse[Response] {
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
