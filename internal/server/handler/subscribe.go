package handler

import (
	"github.com/crypto-grill/app/internal/server/request"
	"go.uber.org/zap"
	"net/http"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	req, err := request.NewSubscribe(r)
	if err != nil {
		zap.S().Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	panic(req)
	// 1. Save to DB
	// 2. Run subscribe_user for channel author
}
