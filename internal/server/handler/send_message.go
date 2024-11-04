package handler

import (
	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/helpers"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func SendMessage(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewSendMessageRequest(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	subscribers, err := ctx.Subscribers(r).New().FilterByChannelID(req.ChannelID).Select()
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdAt := time.Now().UTC()
	msg := data.Message{
		ID:        helpers.RandomID(),
		ChannelID: req.ChannelID,
		Message:   req.Message,
		CreatedAt: &createdAt,
	}
	if err := ctx.Messages(r).New().Save(msg); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 2. Run receive message for subscribers
	panic(subscribers)
}
