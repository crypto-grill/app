package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/handler/incoming"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	postgres2 "github.com/crypto-grill/app/internal/data/postgres"
	"github.com/crypto-grill/app/internal/infrastructure/postgres"
	"github.com/crypto-grill/app/internal/server/handler"
	"github.com/go-chi/chi"
	middleware2 "github.com/go-chi/chi/v5/middleware"

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

	db, err := postgres.New(cfg.Storage.Endpoint)
	if err != nil {
		return nil, err
	}

	ctx2 := context.Background()
	host, err := libp2p.New()

	if err != nil {
		return nil, err
	}
	log.Println("Host created. ID:", host.ID())

	pubsub, err := pubsub.NewGossipSub(ctx2, host)
	if err != nil {
		return nil, err
	}

	r.Use(
		ape.CtxMiddleware(
			ctx.SetUsers(postgres2.NewUsers(db)),
			ctx.SetMessages(postgres2.NewMessages(db)),
			ctx.SetSubscribedChannels(postgres2.NewSubscribedChannels(db)),
			ctx.SetChannels(postgres2.NewChannels(db)),
			ctx.SetSubscribers(postgres2.NewSubscribers(db)),
			ctx.SetSubscriptionProofs(postgres2.NewSubscriptionProofs(db)),

			ctx.SetConfig(cfg),
			ctx.SetHost(host),
			ctx.SetPubSub(pubsub),
		),
		middleware2.DefaultLogger,
	)

	r.Route("/", func(r chi.Router) {
		r.Route("/channels", func(r chi.Router) {
			r.Get("/", handler.GetChannels)
			r.Post("/add", incoming.AddChannel)
			r.Post("/create", handler.CreateChannel)
			r.Post("/subscribe", handler.Subscribe)
			r.Post("/subscribe/user", incoming.SubscribeUser) // as sender
		})
		r.Route("/message", func(r chi.Router) {
			r.Post("/receive", incoming.ReceiveMessage)
			r.Post("/send", handler.SendMessage)
			r.Get("/retransmit", incoming.RetransmitMessages)
		})
	})
	return r, nil
}
