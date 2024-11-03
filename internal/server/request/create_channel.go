package request

import (
	"encoding/json"
	"net/http"
)

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
