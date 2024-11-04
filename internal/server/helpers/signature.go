package helpers

import (
	"encoding/json"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/utils"
	"math/big"
	"time"
)

type SignatureMessage struct {
	UserID    int64     `json:"user_id"`
	ChannelID int64     `json:"channel_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ProofMessage struct {
	SignatureMessage  `json:"signature_message"`
	babyjub.Signature `json:"signature"`
}

func Sign(msg *SignatureMessage, k *babyjub.PrivateKey) (*ProofMessage, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	signature := k.SignPoseidon(utils.SetBigIntFromLEBytes(new(big.Int), bytes))
	return &ProofMessage{
		SignatureMessage: *msg,
		Signature:        *signature,
	}, nil
}

func Verify(proof *ProofMessage, pk *babyjub.PublicKey) (bool, error) {
	if proof.ExpiresAt.Before(time.Now().UTC()) {
		return false, nil
	}
	bytes, err := json.Marshal(proof.SignatureMessage)
	if err != nil {
		return false, err
	}
	msg := utils.SetBigIntFromLEBytes(new(big.Int), bytes)
	valid := pk.VerifyPoseidon(msg, &proof.Signature)
	return valid, nil
}
