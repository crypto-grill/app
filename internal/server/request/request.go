package request

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/crypto-grill/app/internal/server/helpers"
)

type AddChannelRequest struct {
	ID        int64     `json:"id"`
	SenderID  int64     `json:"sender_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAddChannel(r *http.Request) (*AddChannelRequest, error) {
	request := new(AddChannelRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

type CreateChannelRequest struct {
	SenderID int64  `json:"sender_id"`
	Name     string `json:"name"`
}

func NewCreateChannel(r *http.Request) (*CreateChannelRequest, error) {
	request := new(CreateChannelRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

type ReceiveMessageRequest struct {
	ID        int64     `json:"id"`
	ChannelID int64     `json:"channel_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func NewReceiveMessage(r *http.Request) (*ReceiveMessageRequest, error) {
	request := new(ReceiveMessageRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

type RetransmitMessageRequest struct {
	helpers.ProofMessage `json:"proof_message"`
	After                time.Time `json:"after"`
}

func NewRetransmitMessage(r *http.Request) (*RetransmitMessageRequest, error) {
	request := new(RetransmitMessageRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

type SendMessageRequest struct {
	Message   string `json:"message"`
	ChannelID int64  `json:"channel_id"`
}

func NewSendMessageRequest(r *http.Request) (*SendMessageRequest, error) {
	request := new(SendMessageRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

type SubscribeUserRequest struct {
	UserID    int64 `json:"user_id"`
	ChannelID int64 `json:"channel_id"`
}

func NewSubscribeUser(r *http.Request) (*SubscribeUserRequest, error) {
	request := new(SubscribeUserRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

type SubscribeRequest struct {
	UserID    int64 `json:"user_id"`
	ChannelID int64 `json:"channel_id"`
}

func NewSubscribe(r *http.Request) (*SubscribeRequest, error) {
	request := new(SubscribeRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}
