package repository

import (
	"context"
	"errors"
	"slices"

	"github.com/gera9/go-blog/internal/users"
	"github.com/gera9/go-blog/pkg/utils"
	"github.com/google/uuid"
)

type inMemoryRepo struct {
	users []users.User
}

func NewInMemoryRepository() *inMemoryRepo {
	return &inMemoryRepo{}
}

func (r *inMemoryRepo) Create(ctx context.Context, user users.User) (uuid.UUID, error) {
	id := uuid.New()

	user.Id = id

	r.users = append(r.users, user)

	return id, nil
}

func (r *inMemoryRepo) List(ctx context.Context, q users.QueryList) ([]users.User, int, error) {
	return r.users, len(r.users), nil
}

func (r *inMemoryRepo) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *inMemoryRepo) UpdateById(ctx context.Context, id uuid.UUID, user users.User) error {
	for i, v := range r.users {
		if v.Id != id {
			continue
		}

		err := utils.PatchModel(user, v)
		if err != nil {
			return err
		}

		r.users[i] = v

		return nil
	}
	return errors.New("not found")
}

func (r *inMemoryRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	i := 0
	for ; i < len(r.users); i++ {
		if r.users[i].Id == id {
			break
		}
		i++
	}

	r.users = slices.Delete(r.users, i, i+1)

	return nil
}
