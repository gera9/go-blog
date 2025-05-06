package users

import (
	"context"

	"github.com/gera9/go-blog/internal/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (uuid.UUID, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
	List() ([]models.User, int, error)
	UpdateById(ctx context.Context, id uuid.UUID, user models.User) error
}

type service struct {
	userRepo UserRepository
}

func NewService(userRepo UserRepository) *service {
	return &service{userRepo}
}

func (s service) Create(ctx context.Context, user models.User) (uuid.UUID, error) {
	return s.userRepo.Create(ctx, user)
}

func (s service) List() ([]models.User, int, error) {
	return s.userRepo.List()
}

func (s service) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.userRepo.GetById(ctx, id)
}

func (s service) UpdateById(ctx context.Context, id uuid.UUID, user models.User) error {
	return s.userRepo.UpdateById(ctx, id, user)
}

func (s service) DeleteById(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.DeleteById(ctx, id)
}
