package handler

import (
	"github.com/crypto-grill/app/internal/server/ctx"
	"go.uber.org/zap"
	"net/http"
)

func GetMessagesForSubscribedChannels(w http.ResponseWriter, r *http.Request) {
	channels, err := ctx.SubscribedChannels(r).New().SelectChannelIDs()
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	proofs, err := ctx.SubscriptionProofs(r).New().Unexpired().InChannels(channels).Select()
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ips, err := ctx.Users(r).New().GetIPsForChannels(channels)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	/*
		TODO:
		1. Get latest message timestamp for each channel
		2. Request from author
		3. Request from subscribers
	*/

	panic(proofs)
	panic(ips)
}
