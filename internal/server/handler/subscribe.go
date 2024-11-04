package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewSubscribe(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	subscriber := data.Subscriber{
		UserID:    req.UserID,
		ChannelID: req.ChannelID,
	}
	if err := ctx.Subscribers(r).New().Save(subscriber); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get sender id user for channel id

	sender_ip, err := ctx.Users(r).New().GetIPsForChannels([]int64{req.ChannelID})
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url := sender_ip[0] + "channels/subscribe/user"

	jsonData, err := json.Marshal(subscriber)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// TODO maybe delete later (for now)
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Received non-200 response: %d", response.StatusCode)
		return
	}

	subscribeChannel := data.SubscribedChannel{
		ID:        0,
		ChannelID: req.ChannelID,
	}
	if err := ctx.SubscribedChannels(r).New().Save(subscribeChannel); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
