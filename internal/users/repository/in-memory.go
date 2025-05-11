package repository

import (
	"context"
	"errors"
	"sync"

	"slices"

	"github.com/gera9/go-blog/internal/users"
	"github.com/google/uuid"
)

type inMemoryRepo struct {
	mu    sync.RWMutex
	users []users.User
}

func NewInMemoryRepository() *inMemoryRepo {
	return &inMemoryRepo{sync.RWMutex{}, inMemoryUsers}
}

func (r *inMemoryRepo) Create(ctx context.Context, user users.User) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.Id = uuid.New()
	r.users = append(r.users, user)

	return user.Id, nil
}

func (r *inMemoryRepo) List(ctx context.Context, q users.QueryList) ([]users.User, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// For now, we ignore filtering/pagination in QueryList
	return r.users, len(r.users), nil
}

func (r *inMemoryRepo) GetById(ctx context.Context, id uuid.UUID) (*users.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Id == id {
			// return a copy to avoid accidental modification
			u := user
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *inMemoryRepo) UpdateById(ctx context.Context, id uuid.UUID, user users.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, user := range r.users {
		if user.Id == id {
			user.Id = id // ensure ID isn't changed
			r.users[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *inMemoryRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, user := range r.users {
		if user.Id == id {
			r.users = slices.Delete(r.users, i, i+1)
			return nil
		}
	}
	return errors.New("user not found")
}
