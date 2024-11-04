package incoming

import (
	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
	"net/http"
)

func ReceiveMessage(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewReceiveMessage(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	msg := data.Message{
		ID:        req.ID,
		ChannelID: req.ChannelID,
		Message:   req.Message,
		CreatedAt: &req.CreatedAt,
	}
	if err := ctx.Messages(r).New().Save(msg); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
