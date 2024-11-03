package request

import (
	"encoding/json"
	"net/http"
)

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
