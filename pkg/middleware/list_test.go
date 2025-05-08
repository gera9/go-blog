package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_List(t *testing.T) {
	type want struct {
		statusCode int
		limit      int
		offset     int
	}
	tests := []struct {
		name string
		w    http.ResponseWriter
		r    *http.Request
		want want
	}{
		{
			name: "Test List middleware",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/api/v1/users?limit=10&offset=0", nil),
			want: want{
				statusCode: http.StatusOK,
				limit:      10,
				offset:     0,
			},
		},
		{
			name: "Test List middleware with default values",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/api/v1/users", nil),
			want: want{
				statusCode: http.StatusOK,
				limit:      10,
				offset:     0,
			},
		},
		{
			name: "Test List middleware with invalid limit",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/api/v1/users?limit=invalid&offset=0", nil),
			want: want{
				statusCode: http.StatusBadRequest,
				limit:      0,
				offset:     0,
			},
		},
		{
			name: "Test List middleware with invalid offset",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/api/v1/users?limit=10&offset=invalid", nil),
			want: want{
				statusCode: http.StatusBadRequest,
				limit:      10,
				offset:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := MiddlewareManager{}

			nextHanlder := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()

				limit := ctx.Value(ContextKeyLimit)
				if !assert.Equal(t, tt.want.limit, limit) {
					return
				}

				offset := ctx.Value(ContextKeyOffset)
				if !assert.Equal(t, tt.want.offset, offset) {
					return
				}

				w.WriteHeader(http.StatusOK)
			})

			mm.List(nextHanlder).ServeHTTP(tt.w, tt.r)

			rr := tt.w.(*httptest.ResponseRecorder)

			if tt.want.statusCode != rr.Code {
				assert.Equal(t, tt.want.statusCode, rr.Code)
				return
			}
		})
	}
}
