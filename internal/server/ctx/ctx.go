package ctx

import (
	"context"
	"net/http"

	"github.com/crypto-grill/app/internal/config"
)

type ctxKey int

const (
	cfgCtxKey ctxKey = iota + 1
)

func SetConfig(cfg *config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, cfgCtxKey, cfg)
	}
}

func Config(r *http.Request) *config.Config {
	return r.Context().Value(cfgCtxKey).(*config.Config)
}
