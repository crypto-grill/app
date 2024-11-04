package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/config"
	"github.com/crypto-grill/app/internal/server/helpers"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
)

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewCreateChannel(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdAt := time.Now().UTC()
	channel := data.Channel{
		ID:        helpers.RandomID(),
		SenderID:  req.SenderID,
		Name:      req.Name,
		CreatedAt: &createdAt,
	}

	fmt.Printf("Received AddChannel request: ID=%d, SenderID=%d, Name=%s, Time=%s", channel.ID, channel.SenderID, channel.Name, channel.CreatedAt)

	if err := ctx.Channels(r).New().Save(channel); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	users := config.GetUsersWithoutPort(ctx.Config(r).Delivery.HTTP.BindPort)

	for i := 0; i < len(users); i++ {
		url := users[i] + "/channels/add"

		jsonData, err := json.Marshal(channel)
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
	// TODO: Gossip new channel instead
}
