package handler

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/data"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/request"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/utils"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"go.uber.org/zap"
)

func respondeWithProof(w http.ResponseWriter, r *http.Request, status int, subscriber *data.Subscriber) {
	var k babyjub.PrivateKey

	_, err := hex.Decode(k[:], []byte(ctx.SecretKey(r)))
	if err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
		return
	}

	expiration := time.Now().UTC().Add(30 * 24 * time.Hour).Format(time.RFC3339)
	msg := fmt.Sprintf("User Id: %d\nChannel Id: %d\nexp: %s", subscriber.UserID, subscriber.ChannelID, expiration)

	signature := k.SignPoseidon(utils.SetBigIntFromLEBytes(new(big.Int), []byte(msg)))

	response := map[string]interface{}{
		"message": msg,
		"proof": map[string]string{
			"R8_X": signature.R8.X.String(),
			"R8_Y": signature.R8.Y.String(),
			"S":    signature.S.String(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		zap.S().Error("Failed to encode JSON response: ", err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
	}
}

func SubscribeUser(w http.ResponseWriter, r *http.Request) {
	// 1. Save subscriber to local db
	// 2. Return proof of subscription.

	req, err := request.NewSubscribeUser(r)
	if err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
		return
	}

	subscriber := data.Subscriber{
		UserID:    req.UserID,
		ChannelID: req.ChannelID,
	}

	if err := ctx.Subscribers(r).New().Save(subscriber); err != nil {
		zap.S().Error(err)
		errObjects := problems.BadRequest(err)
		ape.RenderErr(w, errObjects...)
	}

	respondeWithProof(w, r, http.StatusOK, &subscriber)

	//w.WriteHeader(http.StatusOK)
}
