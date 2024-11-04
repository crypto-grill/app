package handler

import (
	"encoding/json"
	"net/http"

	"github.com/crypto-grill/app/internal/server/ctx"
	"go.uber.org/zap"
)

func GetChannels(w http.ResponseWriter, r *http.Request) {
	channels, err := ctx.Channels(r).New().Select()
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(channels); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: reverse gossip transport protocol (RGTP)
	// 2. GetChannels from other nodes if not and save to db
}
