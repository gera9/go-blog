package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gera9/go-blog/config"
	"github.com/gera9/go-blog/internal/users"
	"github.com/gera9/go-blog/pkg/logger"
	"github.com/gera9/go-blog/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.SetUpZap()
	defer log.Sync()

	postgresConn, err := postgres.GetPostgresConn(cfg.Postgres.Connstr)
	if err != nil {
		log.Fatal("failed to connect to postgres", zap.Error(err))
	}
	defer postgresConn.Close()

	userRepo := users.NewRepository(postgresConn)
	userSvc := users.NewService(userRepo)
	userCtller := users.NewHttpController(userSvc)

	r := chi.NewMux()

	r.Mount("/users", userCtller.Routes())

	log.Sugar().Infof("Server running on port: %d", cfg.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), r)
}
