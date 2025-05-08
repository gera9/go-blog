package users

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, user User) (uuid.UUID, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	List(ctx context.Context, q QueryList) ([]User, int, error)
	UpdateById(ctx context.Context, id uuid.UUID, user User) error
}

type Service interface {
	Create(ctx context.Context, user User) (uuid.UUID, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	List(ctx context.Context, q QueryList) ([]User, int, error)
	UpdateById(ctx context.Context, id uuid.UUID, user User) error
}
