package delivery

import (
	"errors"
	"net/http"

	shareddtos "github.com/gera9/go-blog/internal/shared-models/dtos"
	"github.com/gera9/go-blog/internal/users"
	"github.com/gera9/go-blog/pkg/middleware"
	"github.com/gera9/go-blog/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type httpController struct {
	userSvc users.Service
}

func NewHttpController(userSvc users.Service) *httpController {
	return &httpController{userSvc}
}

func (c httpController) Routes(mm *middleware.MiddlewareManager) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", c.Create)
	r.With(mm.List).Get("/", c.List)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", c.GetById)
		r.Patch("/", c.UpdateById)
		r.Delete("/", c.DeleteById)
	})

	return r
}

// CreateUser godoc
// @Summary Create new user
// @Description Create new user
// @Tags Users
// @Accept json
// @Produce json
// @Param paquete body CreatePayload true "Create user payload"
// @Success 201 {object} shareddtos.IdResponse
// @Failure 400 {object} shareddtos.ErrResponse
// @Failure 500 {object} shareddtos.ErrResponse
// @Router /users [post]
func (c httpController) Create(w http.ResponseWriter, r *http.Request) {
	payload := CreatePayload{}
	err := render.Bind(r, &payload)
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(err, errors.New("invalid body")))
		return
	}

	if !payload.IsValid() {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid json")))
		return
	}

	id, err := c.userSvc.Create(r.Context(), payload.ToModel())
	if err != nil {
		render.Render(w, r, shareddtos.NewInternalServerErr(err))
		return
	}

	render.JSON(w, r, shareddtos.IdResponse{Id: id.String()})
	w.WriteHeader(http.StatusOK)
}

// ListUsers godoc
// @Summary List users
// @Description List users
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} ListResponse
// @Failure 500 {object} shareddtos.ErrResponse
// @Router /users [get]
func (c httpController) List(w http.ResponseWriter, r *http.Request) {
	q := users.QueryList{
		Limit:  r.Context().Value(middleware.ContextKeyLimit).(int),
		Offset: r.Context().Value(middleware.ContextKeyOffset).(int),
	}

	userss, total, err := c.userSvc.List(r.Context(), q)
	if err != nil {
		render.Render(w, r, shareddtos.NewInternalServerErr(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, ToListResponse(total, userss))
}

// GetById godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} Response
// @Failure 400 {object} shareddtos.ErrResponse
// @Failure 500 {object} shareddtos.ErrResponse
// @Router /users/{id} [get]
func (c httpController) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.IdURLParamToUUID(r, "id")
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(err, errors.New("invalid id")))
		return
	}

	user, err := c.userSvc.GetById(r.Context(), id)
	if err != nil {
		render.Render(w, r, shareddtos.NewInternalServerErr(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, ToResponse(*user))
}

// UpdateUserById godoc
// @Summary Update user by ID
// @Description Update user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdatePayload true "Update user payload"
// @Success 204
// @Failure 400 {object} shareddtos.ErrResponse
// @Failure 500 {object} shareddtos.ErrResponse
// @Router /users/{id} [patch]
func (c httpController) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.IdURLParamToUUID(r, "id")
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(err, errors.New("invalid id")))
		return
	}

	payload := UpdatePayload{}
	err = render.Bind(r, &payload)
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(err, errors.New("invalid body")))
		return
	}

	if !payload.IsValid() {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid json")))
		return
	}

	err = c.userSvc.UpdateById(r.Context(), id, payload.ToModel())
	if err != nil {
		render.Render(w, r, shareddtos.NewInternalServerErr(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteById godoc
// @Summary Delete user by ID
// @Description Delete user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 400 {object} shareddtos.ErrResponse
// @Failure 500 {object} shareddtos.ErrResponse
// @Router /users/{id} [delete]
func (c httpController) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.IdURLParamToUUID(r, "id")
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(err, errors.New("invalid id")))
		return
	}

	err = c.userSvc.DeleteById(r.Context(), id)
	if err != nil {
		render.Render(w, r, shareddtos.NewInternalServerErr(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
