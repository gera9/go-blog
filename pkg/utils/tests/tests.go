package testutils

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	ExampleId = "24318026-421d-4aaf-bbf1-b2f7c4715597"
)

type SetUpMock func(m *mock.Mock)

func SetUpReturnMock(methodName string, arguments, returnArguments []any) SetUpMock {
	return func(m *mock.Mock) {
		m.On(methodName, arguments...).
			Return(returnArguments...)
	}
}

func AssertJsonBodyAndStatusCodeResponse(t *testing.T, w http.ResponseWriter, expectedStatusCode int, expectedResponse string) {
	rr := w.(*httptest.ResponseRecorder)

	body := rr.Body.String()

	if expectedResponse == "" {
		if body != "" {
			t.Errorf("body: %s != expectedResponse: %s", body, expectedResponse)
			return
		} else {
			return
		}
	}

	if rr.Code != expectedStatusCode {
		t.Errorf("%d != %d: %s", expectedStatusCode, rr.Code, body)
		return
	}

	assert.JSONEq(t, expectedResponse, body)
}

func SetRequestBodyFromFile(r *http.Request, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, f); err != nil {
		return err
	}

	r.Body = io.NopCloser(buf)

	return f.Close()
}

func SetRequestJsonFromFile(r *http.Request, filename string) error {
	err := SetRequestBodyFromFile(r, filename)
	if err != nil {
		return err
	}

	r.Header.Set("Content-Type", "application/json")

	return nil
}

func AddURLParam(r *http.Request, key, value string) {
	ctx := chi.RouteContext(r.Context())

	if ctx == nil {
		ctx = chi.NewRouteContext()
	}

	ctx.URLParams.Add(key, value)

	*r = *r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
}
