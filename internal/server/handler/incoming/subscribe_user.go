package incoming

import (
	"encoding/hex"
	"encoding/json"
	"github.com/crypto-grill/app/internal/server/helpers"
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"go.uber.org/zap"
)

// We hardcode subscriptions to last for `month` (30 days).
const month = 30 * 24 * time.Hour

func SubscribeUser(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewSubscribeUser(r)
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
	proof, err := formProof(ctx.Config(r).SecretKey, &subscriber)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(proof); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func formProof(sk string, subscriber *data.Subscriber) (*helpers.ProofMessage, error) {
	var k babyjub.PrivateKey
	_, err := hex.Decode(k[:], []byte(sk))
	if err != nil {
		return nil, err
	}
	msg := helpers.SignatureMessage{
		UserID:    subscriber.UserID,
		ChannelID: subscriber.ChannelID,
		ExpiresAt: time.Now().Add(month).UTC(),
	}
	return helpers.Sign(&msg, &k)
}
