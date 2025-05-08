package repository

import (
	"context"
	"database/sql"

	"github.com/gera9/go-blog/internal/users"
	"github.com/gera9/go-blog/pkg/postgres"
	"github.com/google/uuid"
)

type postgresRepo struct {
	*sql.DB
}

func NewPostgresRepository(postgres *postgres.PostgresConn) *postgresRepo {
	return &postgresRepo{postgres.DB}
}

func (r postgresRepo) Create(ctx context.Context, user users.User) (uuid.UUID, error) {
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

func (r postgresRepo) List(ctx context.Context, q users.QueryList) ([]users.User, int, error) {
	return nil, 0, nil
}

func (r postgresRepo) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	return nil, nil
}

func (r postgresRepo) UpdateById(ctx context.Context, id uuid.UUID, user users.User) error {
	return nil
}

func (r postgresRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	return nil
}
