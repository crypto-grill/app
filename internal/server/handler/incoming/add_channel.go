package incoming

import (
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
)

func AddChannel(w http.ResponseWriter, r *http.Request) {
	return
	req, err := request.NewAddChannel(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Printf("Received AddChannel request: ID=%d, SenderID=%d, Name=%s, Time=%s", req.ID, req.SenderID, req.Name, req.CreatedAt)

	time := time.Now().UTC()
	channel := data.Channel{
		ID:        1,
		SenderID:  req.SenderID,
		Name:      req.Name,
		CreatedAt: &time,
	}

	if err := ctx.Channels(r).New().Save(channel); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

	// TODO: Gossip channel
}
