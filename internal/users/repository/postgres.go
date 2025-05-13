package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gera9/go-blog/internal/users"
	"github.com/gera9/go-blog/pkg/postgres"
	"github.com/google/uuid"
)

type postgresRepo struct {
	*sql.DB
	tableName string
}

func NewPostgresRepository(postgres *postgres.PostgresConn) *postgresRepo {
	return &postgresRepo{postgres.DB, "users"}
}

func (r postgresRepo) Create(ctx context.Context, user users.User) (uuid.UUID, error) {
	insertedId := uuid.UUID{}

	err := r.DB.QueryRow(Create,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Birthdate,
		user.PasswordHash,
	).Scan(&insertedId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Create: %v", err)
	}

	return insertedId, nil
}

func (r postgresRepo) List(ctx context.Context, q users.QueryList) ([]users.User, int, error) {
	rows, err := r.DB.Query(List, q.Limit, q.Offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var res []users.User
	for rows.Next() {
		var user users.User

		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.Birthdate,
			&user.PasswordHash,
		); err != nil {
			return nil, 0, err
		}

		res = append(res, user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return res, len(res), nil
}

func (r postgresRepo) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	res := &users.User{}
	err := r.DB.QueryRow(GetById, id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.Email,
		&res.Birthdate,
		&res.PasswordHash,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("GetById %d: no such user", id)
	}

	return res, nil
}

func (r postgresRepo) UpdateById(ctx context.Context, id uuid.UUID, user users.User) error {
	return nil
}

func (r postgresRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	_, err := r.DB.Exec(DeleteById, id)
	if err != nil {
		return fmt.Errorf("DeleteById %d: %v", id, err)
	}

	return nil
}
