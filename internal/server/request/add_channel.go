package request

import (
	"encoding/json"
	"net/http"
)

type AddChannelRequest struct {
	ID       int64  `json:"id"`
	SenderID int64  `json:"sender_id"`
	Name     string `json:"name"`
}

func NewAddChannel(r *http.Request) (*AddChannelRequest, error) {
	request := new(AddChannelRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}
