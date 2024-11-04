package handler

import (
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/server/helpers"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
)

// CreateChannel is also an incoming handler
func CreateChannel(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewCreateChannel(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdAt := time.Now().UTC()
	channel := data.Channel{
		ID:        helpers.RandomID(),
		SenderID:  req.SenderID,
		Name:      req.Name,
		CreatedAt: &createdAt,
	}
	if err := ctx.Channels(r).New().Save(channel); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	// TODO: Gossip new channel
}
