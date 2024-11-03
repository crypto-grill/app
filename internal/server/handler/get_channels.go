package handler

import "net/http"

func GetChannels(w http.ResponseWriter, r *http.Request) {
	// 1. Check record in local DB and return it if exists
	// 2. GetChannels from other nodes if not and save to db
}
