package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/gera9/go-blog/internal/users"
	"github.com/gera9/go-blog/internal/users/mocks"
	testutils "github.com/gera9/go-blog/pkg/utils/tests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func Test_service_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user users.User
	}
	tests := []struct {
		name      string
		setUpMock testutils.SetUpMock
		args      args
		want      uuid.UUID
		wantErr   bool
	}{
		{
			name: "Create ok",
			setUpMock: testutils.SetUpReturnMock(
				"Create",
				[]any{mock.Anything, mock.Anything},
				[]any{uuid.Nil, nil},
			),
			want:    uuid.Nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.MockRepository)

			tt.setUpMock(repoMock)

			s := service{
				userRepo: repoMock,
			}

			got, err := s.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_List(t *testing.T) {
	type args struct {
		ctx context.Context
		q   users.QueryList
	}
	tests := []struct {
		name      string
		setUpMock testutils.SetUpMock
		args      args
		want      []users.User
		want1     int
		wantErr   bool
	}{
		{
			name: "List ok",
			setUpMock: testutils.SetUpReturnMock(
				"List",
				[]any{mock.Anything, mock.Anything},
				[]any{
					[]users.User{}, 1, nil,
				},
			),
			want:    []users.User{},
			want1:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.MockRepository)

			tt.setUpMock(repoMock)

			s := service{
				userRepo: repoMock,
			}

			got, got1, err := s.List(tt.args.ctx, tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.List() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("service.List() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_service_GetById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name      string
		setUpMock testutils.SetUpMock
		args      args
		want      *users.User
		wantErr   bool
	}{
		{
			name: "GetById ok",
			setUpMock: testutils.SetUpReturnMock(
				"GetById",
				[]any{mock.Anything, mock.Anything},
				[]any{&users.User{}, nil},
			),
			want:    &users.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.MockRepository)

			tt.setUpMock(repoMock)

			s := service{
				userRepo: repoMock,
			}

			got, err := s.GetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateById(t *testing.T) {
	type args struct {
		ctx  context.Context
		id   uuid.UUID
		user users.User
	}
	tests := []struct {
		name      string
		setUpMock testutils.SetUpMock
		args      args
		wantErr   bool
	}{
		{
			name: "UpdateById ok",
			setUpMock: testutils.SetUpReturnMock(
				"UpdateById",
				[]any{mock.Anything, mock.Anything, mock.Anything},
				[]any{nil},
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.MockRepository)

			tt.setUpMock(repoMock)

			s := service{
				userRepo: repoMock,
			}

			if err := s.UpdateById(tt.args.ctx, tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name      string
		setUpMock testutils.SetUpMock
		args      args
		wantErr   bool
	}{
		{
			name: "DeleteById ok",
			setUpMock: testutils.SetUpReturnMock(
				"DeleteById",
				[]any{mock.Anything, mock.Anything},
				[]any{nil},
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.MockRepository)

			tt.setUpMock(repoMock)

			s := service{
				userRepo: repoMock,
			}

			if err := s.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
