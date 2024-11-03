package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"go.uber.org/zap"
)

func CreateChannels(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewCreateChannel(r)
	if err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
		return
	}

	// Random id generate for channel. Maybe bad idea
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	channel := data.Channel{
		ID:       random.Int63(),
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
