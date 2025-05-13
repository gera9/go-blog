package delivery

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gera9/go-blog/internal/users"
	"github.com/gera9/go-blog/internal/users/mocks"
	"github.com/gera9/go-blog/pkg/middleware"
	testutils "github.com/gera9/go-blog/pkg/utils/tests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mm middleware.MiddlewareManager

func Test_httpController_Create(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name               string
		args               args
		mockSetup          testutils.SetUpMock
		expectedStatusCode int
		expectedResponse   string
		bodyPath           string
	}{
		{
			name: "Create ok",
			mockSetup: testutils.SetUpReturnMock(
				"Create",
				[]any{mock.Anything, mock.Anything},
				[]any{uuid.Nil, nil},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/users", nil),
			},
			bodyPath:           "./testdata/create/ok-body.json",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"id":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name: "Create service error",
			mockSetup: testutils.SetUpReturnMock(
				"Create",
				[]any{mock.Anything, mock.Anything},
				[]any{nil, assert.AnError},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/users", nil),
			},
			bodyPath:           "./testdata/create/ok-body.json",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"dev_message":"assert.AnError general error for testing", "status_code":500, "user_message":"assert.AnError general error for testing"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := new(mocks.MockService)

			tt.mockSetup(serviceMock)

			c := httpController{
				userSvc: serviceMock,
			}

			err := testutils.SetRequestJsonFromFile(tt.args.r, tt.bodyPath)
			if err != nil {
				t.Error(err)
				return
			}

			c.Create(tt.args.w, tt.args.r)

			testutils.AssertJsonBodyAndStatusCodeResponse(t, tt.args.w, tt.expectedStatusCode, tt.expectedResponse)
		})
	}
}

func Test_httpController_List(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name               string
		args               args
		mockSetup          testutils.SetUpMock
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "List ok",
			mockSetup: testutils.SetUpReturnMock(
				"List",
				[]any{mock.Anything, mock.Anything},
				[]any{
					[]users.User{
						{
							Id:           uuid.Nil,
							FirstName:    "Alice",
							LastName:     "Smith",
							Username:     "alice123",
							Email:        "email@email.com",
							PasswordHash: "passAlice!",
							Birthdate:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
						},
						{
							Id:           uuid.Nil,
							FirstName:    "Bob",
							LastName:     "Johnson",
							Username:     "bobbyj",
							Email:        "email@email.com",
							PasswordHash: "bobPass2024",
							Birthdate:    time.Date(1985, 8, 20, 0, 0, 0, 0, time.UTC),
						},
					}, 2, nil,
				},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/users?limit=10&offset=0", nil),
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"total":2,"items":[{"id":"00000000-0000-0000-0000-000000000000","first_name":"Alice","email":"email@email.com","last_name":"Smith","username":"alice123","birthdate":"1990-05-10T00:00:00Z"},{"id":"00000000-0000-0000-0000-000000000000","first_name":"Bob","email":"email@email.com","last_name":"Johnson","username":"bobbyj","birthdate":"1985-08-20T00:00:00Z"}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := new(mocks.MockService)

			tt.mockSetup(serviceMock)

			c := httpController{
				userSvc: serviceMock,
			}

			mm.List(http.HandlerFunc(c.List)).ServeHTTP(tt.args.w, tt.args.r)

			testutils.AssertJsonBodyAndStatusCodeResponse(t, tt.args.w, tt.expectedStatusCode, tt.expectedResponse)
		})
	}
}

func Test_httpController_GetById(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name               string
		args               args
		mockSetup          testutils.SetUpMock
		id                 string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "GetById ok",
			mockSetup: testutils.SetUpReturnMock(
				"GetById",
				[]any{mock.Anything, mock.Anything},
				[]any{
					&users.User{

						Id:           uuid.Nil,
						FirstName:    "Alice",
						LastName:     "Smith",
						Username:     "alice123",
						Email:        "email@email.com",
						PasswordHash: "passAlice!",
						Birthdate:    time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
					},
					nil,
				},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/users", nil),
			},
			id:                 testutils.ExampleId,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"id":"00000000-0000-0000-0000-000000000000","first_name":"Alice","email":"email@email.com","last_name":"Smith","username":"alice123","birthdate":"1990-05-10T00:00:00Z"}`,
		},
		{
			name: "GetById service error",
			mockSetup: testutils.SetUpReturnMock(
				"GetById",
				[]any{mock.Anything, mock.Anything},
				[]any{
					nil,
					assert.AnError,
				},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/users", nil),
			},
			id:                 testutils.ExampleId,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"dev_message":"assert.AnError general error for testing", "status_code":500, "user_message":"assert.AnError general error for testing"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := new(mocks.MockService)

			tt.mockSetup(serviceMock)

			c := httpController{
				userSvc: serviceMock,
			}

			testutils.AddURLParam(tt.args.r, "id", tt.id)

			c.GetById(tt.args.w, tt.args.r)

			testutils.AssertJsonBodyAndStatusCodeResponse(t, tt.args.w, tt.expectedStatusCode, tt.expectedResponse)
		})
	}
}

func Test_httpController_UpdateById(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name               string
		args               args
		mockSetup          testutils.SetUpMock
		id                 string
		expectedStatusCode int
		expectedResponse   string
		bodyPath           string
	}{
		{
			name: "UpdateById ok",
			mockSetup: testutils.SetUpReturnMock(
				"UpdateById",
				[]any{mock.Anything, mock.Anything, mock.Anything},
				[]any{nil},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPatch, "/users", nil),
			},
			id:                 testutils.ExampleId,
			expectedStatusCode: http.StatusNoContent,
			bodyPath:           "./testdata/create/ok-body.json",
		},
		{
			name: "UpdateById service error",
			mockSetup: testutils.SetUpReturnMock(
				"UpdateById",
				[]any{mock.Anything, mock.Anything, mock.Anything},
				[]any{assert.AnError},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPatch, "/users", nil),
			},
			id:                 testutils.ExampleId,
			expectedStatusCode: http.StatusInternalServerError,
			bodyPath:           "./testdata/create/ok-body.json",
			expectedResponse:   `{"status_code":500,"dev_message":"assert.AnError general error for testing","user_message":"assert.AnError general error for testing"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutils.SetRequestJsonFromFile(tt.args.r, tt.bodyPath)
			if err != nil {
				t.Error(err)
				return
			}

			serviceMock := new(mocks.MockService)

			tt.mockSetup(serviceMock)

			c := httpController{
				userSvc: serviceMock,
			}

			testutils.AddURLParam(tt.args.r, "id", tt.id)

			c.UpdateById(tt.args.w, tt.args.r)

			testutils.AssertJsonBodyAndStatusCodeResponse(t, tt.args.w, tt.expectedStatusCode, tt.expectedResponse)
		})
	}
}

func Test_httpController_DeleteById(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name               string
		args               args
		mockSetup          testutils.SetUpMock
		id                 string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "DeleteById ok",
			mockSetup: testutils.SetUpReturnMock(
				"DeleteById",
				[]any{mock.Anything, mock.Anything},
				[]any{nil},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodDelete, "/users", nil),
			},
			id:                 testutils.ExampleId,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "DeleteById service error",
			mockSetup: testutils.SetUpReturnMock(
				"DeleteById",
				[]any{mock.Anything, mock.Anything},
				[]any{assert.AnError},
			),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodDelete, "/users", nil),
			},
			id:                 testutils.ExampleId,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"status_code":500,"dev_message":"assert.AnError general error for testing","user_message":"assert.AnError general error for testing"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceMock := new(mocks.MockService)

			tt.mockSetup(serviceMock)

			c := httpController{
				userSvc: serviceMock,
			}

			testutils.AddURLParam(tt.args.r, "id", tt.id)

			c.DeleteById(tt.args.w, tt.args.r)

			testutils.AssertJsonBodyAndStatusCodeResponse(t, tt.args.w, tt.expectedStatusCode, tt.expectedResponse)
		})
	}
}
