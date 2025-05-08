package middleware

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	shareddtos "github.com/gera9/go-blog/internal/shared-models/dtos"
	"github.com/go-chi/render"
)

var (
	ContextKeyLimit  = &ContextKey{"limit"}
	ContextKeyOffset = &ContextKey{"offset"}
)

// List middleware to handle list requests.
func (mm MiddlewareManager) List(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		limitStr := query.Get("limit")
		if limitStr == "" {
			limitStr = "10"
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid limit")))
			return
		}

		offsetStr := query.Get("offset")
		if offsetStr == "" {
			offsetStr = "0"
		}
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			render.Render(w, r, shareddtos.NewBadRequestErr(errors.New("invalid offset")))
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, ContextKeyLimit, limit)
		ctx = context.WithValue(ctx, ContextKeyOffset, offset)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
