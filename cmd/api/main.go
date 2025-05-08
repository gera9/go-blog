package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gera9/go-blog/config"
	"github.com/gera9/go-blog/internal/users/delivery"
	"github.com/gera9/go-blog/internal/users/repository"
	"github.com/gera9/go-blog/internal/users/service"
	"github.com/gera9/go-blog/pkg/logger"
	"github.com/gera9/go-blog/pkg/middleware"
	"github.com/gera9/go-blog/pkg/postgres"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// @title Go-Blog API
// @version 1.0
// @description Basic API for an API implementation.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic  BasicAuth

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.NewZapLogger(cfg)
	defer log.Sync()

	postgresConn, err := postgres.GetPostgresConn(cfg.Postgres.Connstr)
	if err != nil {
		log.Fatal("failed to connect to postgres", zap.Error(err))
	}
	defer postgresConn.Close()

	userRepo := repository.NewPostgresRepository(postgresConn)
	userSvc := service.NewService(userRepo)
	userCtller := delivery.NewHttpController(userSvc)

	r := chi.NewRouter()
	r.Use(
		chiMiddleware.RequestID,
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
	)

	mm := &middleware.MiddlewareManager{}

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use()
			r.Mount("/users", userCtller.Routes(mm))
		})
	})

	log.Sugar().Infof("Server running on port: %d", cfg.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), r)
}
