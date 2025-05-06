package users

import (
	"context"
	"database/sql"

	"github.com/gera9/go-blog/internal/models"
	"github.com/gera9/go-blog/pkg/postgres"
	"github.com/google/uuid"
)

type repository struct {
	*sql.DB
}

func NewRepository(postgres *postgres.PostgresConn) *repository {
	return &repository{postgres.DB}
}

func (r repository) Create(ctx context.Context, user models.User) (uuid.UUID, error) {
	row := r.DB.QueryRowContext(ctx, `
	INSERT INTO users VALUES ()
	`)

	err := row.Err()
	if err != nil {
		return uuid.UUID{}, err
	}

	id := uuid.UUID{}
	err = row.Scan(&id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (r repository) List() ([]models.User, int, error) {
	return nil, 0, nil
}

func (r repository) GetById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (r repository) UpdateById(ctx context.Context, id uuid.UUID, user models.User) error {
	return nil
}

func (r repository) DeleteById(ctx context.Context, id uuid.UUID) error {
	return nil
}
