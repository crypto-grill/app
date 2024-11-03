package handler

import (
	"encoding/json"
	"net/http"

	"github.com/crypto-grill/app/internal/server/ctx"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"go.uber.org/zap"
)

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		zap.S().Error("Failed to encode JSON response: ", err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
	}
}

func GetChannels(w http.ResponseWriter, r *http.Request) {
	// 1. Check record in local DB and return it if exists
	// 2. GetChannels from other nodes if not and save to db

	channels, err := ctx.Channels(r).New().Get()
	if err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
	}

	if len(channels) > 0 {
		respondWithJSON(w, http.StatusOK, channels)
		return
	}

	// TODO where can we find paths to other users?

	w.WriteHeader(http.StatusOK)
}
