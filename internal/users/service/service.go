package service

import (
	"context"

	"github.com/gera9/go-blog/internal/users"
	"github.com/google/uuid"
)

type service struct {
	userRepo users.Repository
}

func NewService(userRepo users.Repository) *service {
	return &service{userRepo}
}

func (s service) Create(ctx context.Context, user users.User) (uuid.UUID, error) {
	return s.userRepo.Create(ctx, user)
}

func (s service) List(ctx context.Context, q users.QueryList) ([]users.User, int, error) {
	return s.userRepo.List(ctx, q)
}

func (s service) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	return s.userRepo.GetById(ctx, id)
}

func (s service) UpdateById(ctx context.Context, id uuid.UUID, user users.User) error {
	return s.userRepo.UpdateById(ctx, id, user)
}

func (s service) DeleteById(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.DeleteById(ctx, id)
}
