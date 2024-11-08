package incoming

import (
	"fmt"
	"net/http"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
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

	name, err := ctx.Channels(r).New().GetName(msg.ChannelID)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s: %s\n", name, msg.Message)
	w.WriteHeader(http.StatusNoContent)
}
