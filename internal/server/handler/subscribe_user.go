package handler

import "net/http"

func SubscribeUser(w http.ResponseWriter, r *http.Request) {
	// 1. Save subscriber to local db
	// 2. Return proof of subscription.
}
