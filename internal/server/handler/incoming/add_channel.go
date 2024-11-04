package incoming

import (
	"fmt"
	"net/http"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
)

func AddChannel(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewAddChannel(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("Received AddChannel request: ID=%d, SenderID=%d, Name=%s, Time", req.ID, req.SenderID, req.Name, req.CreatedAt)

	channel := data.Channel{
		ID:        req.ID,
		SenderID:  req.SenderID,
		Name:      req.Name,
		CreatedAt: &req.CreatedAt,
	}

	if err := ctx.Channels(r).New().Save(channel); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

	// TODO: Gossip channel
}
