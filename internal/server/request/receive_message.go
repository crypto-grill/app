package request

import (
	"encoding/json"
	"net/http"
	"time"
)

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
