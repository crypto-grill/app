package handler

import (
	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"go.uber.org/zap"
	"net/http"
)

func ReceiveMessage(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewReceiveMessage(r)
	if err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
		return
	}

	msg := data.Message{
		ID:        req.ID,
		ChannelID: req.ChannelID,
		Message:   req.Message,
		CreatedAt: req.CreatedAt,
	}

	if err := ctx.Messages(r).New().Save(msg); err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
	}

	w.WriteHeader(http.StatusOK)
}
