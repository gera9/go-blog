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
// @Summary Creates a new user
// @Description Creates a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param paquete body dtos.CreatePayload true "User payload"
// @Success 201 {object} shareddtos.IdResponse
// @Failure 400 {object} shareddtos.ErrResponse
// @Failure 500 {object} shareddtos.ErrResponse
// @Router /users [post]
func (c httpController) Create(w http.ResponseWriter, r *http.Request) {
	payload := users.CreatePayload{}
	err := render.Bind(r, &payload)
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid body")))
		return
	}

	if !payload.IsValid() {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid json")))
		return
	}

	id, err := c.userSvc.Create(r.Context(), payload.ToModel())
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(err))
		return
	}

	render.JSON(w, r, shareddtos.IdResponse{Id: id.String()})
	w.WriteHeader(http.StatusOK)
}

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
	render.JSON(w, r, users.ToListResponse(total, userss))
}

func (c httpController) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.IdURLParamToUUID(r, "id")
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid id")))
		return
	}

	user, err := c.userSvc.GetById(r.Context(), id)
	if err != nil {
		render.Render(w, r, shareddtos.NewInternalServerErr(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, users.ToResponse(*user))
}

func (c httpController) UpdateById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.IdURLParamToUUID(r, "id")
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid id")))
		return
	}

	payload := users.UpdatePayload{}
	err = render.Bind(r, &payload)
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid body")))
		return
	}

	if !payload.IsValid() {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid json")))
		return
	}

	err = c.userSvc.UpdateById(r.Context(), id, payload.ToModel())
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c httpController) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.IdURLParamToUUID(r, "id")
	if err != nil {
		render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid id")))
		return
	}

	err = c.userSvc.DeleteById(r.Context(), id)
	if err != nil {
		render.Render(w, r, shareddtos.NewInternalServerErr(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
