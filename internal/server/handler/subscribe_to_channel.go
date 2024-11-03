package handler

import "net/http"

func SubscribeToChannel(w http.ResponseWriter, r *http.Request) {
	// 1. Save to DB
	// 2. Run subscribe_user for channel author
}
