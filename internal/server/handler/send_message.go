package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/helpers"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
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

	userIDs := make([]int64, len(subscribers))
	for i, subscriber := range subscribers {
		userIDs[i] = subscriber.UserID
	}

	ips, err := ctx.Users(r).New().GetIPsForSubsriber(userIDs)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(ips); i++ {
		url := ips[i] + "/message/receive"

		jsonData, err := json.Marshal(msg)
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
		}
	}
	// 2. Run receive message for subscribers
	//panic(subscribers)
}
