package handler

import "net/http"

func AddChannel(w http.ResponseWriter, r *http.Request) {
	// 1. Save new channel to DB
	// 2. AddChannel to neighbor nodes
}
