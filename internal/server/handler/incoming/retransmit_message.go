package incoming

import (
	"encoding/json"
	"github.com/crypto-grill/app/internal/server/ctx"
	"github.com/crypto-grill/app/internal/server/helpers"
	"github.com/crypto-grill/app/internal/server/request"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"go.uber.org/zap"
	"net/http"
)

func RetransmitMessages(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewRetransmitMessage(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	stringPK, err := ctx.Users(r).New().GetPubKeyForChannel(req.ChannelID)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pk := &babyjub.PublicKey{}
	if err := pk.UnmarshalText([]byte(stringPK)); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	eligible, err := helpers.Verify(&req.ProofMessage, pk)
	if !eligible || err != nil {
		zap.S().Info(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	messages, err := ctx.Messages(r).New().InChannel(req.ChannelID).After(req.After).Select()
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
