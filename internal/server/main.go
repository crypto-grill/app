package server

import (
	"fmt"
	"github.com/crypto-grill/app/internal/infrastructure/postgres"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/go-chi/chi"
	middleware2 "github.com/go-chi/chi/v5/middleware"
	"net"
	"net/http"

	"github.com/crypto-grill/app/internal/config"
	"gitlab.com/distributed_lab/ape"
)

func Start(cfg *config.Config) error {
	address := fmt.Sprintf(":%d", cfg.Delivery.HTTP.BindPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w", address, err)
	}
	router, err := newRouter(cfg)
	if err != nil {
		return err
	}
	server := &http.Server{
		Handler:     router,
		ReadTimeout: cfg.Delivery.HTTP.RequestTimeout,
	}
	return server.Serve(listener)
}

func newRouter(cfg *config.Config) (chi.Router, error) {
	r := chi.NewRouter()

	_, err := postgres.New(cfg.Storage.Endpoint)
	if err != nil {
		return nil, err
	}
	r.Use(
		ape.CtxMiddleware(

			ctx.SetConfig(cfg),
		),
		middleware2.DefaultLogger,
	)
	r.Route("/", func(r chi.Router) {

	})
	return r, nil
}
