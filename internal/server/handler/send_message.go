package handler

import "net/http"

func SendMessage(w http.ResponseWriter, r *http.Request) {
	// 1. Save to local DB
	// 2. Run receive message for subscribers
}
