package handler

import (
	"net/http"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"go.uber.org/zap"
)

func AddChannel(w http.ResponseWriter, r *http.Request) {
	// 1. Save new channel to DB
	// 2. AddChannel to neighbor nodes

	req, err := request.NewAddChannel(r)
	if err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
		return
	}

	channel := data.Channel{
		ID:       req.ID,
		SenderID: req.SenderID,
		Name:     req.Name,
	}

	if err := ctx.Channels(r).New().Save(channel); err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
	}

	// TODO Gossip protocol
	// 2. Add channel to nearby nodes

	w.WriteHeader(http.StatusOK)
}
