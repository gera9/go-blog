package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gera9/go-blog/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, user models.User) (uuid.UUID, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
	List() ([]models.User, int, error)
	UpdateById(ctx context.Context, id uuid.UUID, user models.User) error
}

type httpController struct {
	userSvc UserService
}

func NewHttpController(userSvc UserService) *httpController {
	return &httpController{userSvc}
}

func (c httpController) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", c.Create)
	r.Get("/", c.List)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", c.GetById)
		r.Patch("/", c.UpdateById)
		r.Delete("/", c.DeleteById)
	})

	return r
}

func (c httpController) Create(w http.ResponseWriter, r *http.Request) {
}

func (c httpController) List(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Hello, from users 2!",
	})

	w.WriteHeader(http.StatusOK)
}

func (c httpController) GetById(w http.ResponseWriter, r *http.Request) {
}

func (c httpController) UpdateById(w http.ResponseWriter, r *http.Request) {
}

func (c httpController) DeleteById(w http.ResponseWriter, r *http.Request) {
}
