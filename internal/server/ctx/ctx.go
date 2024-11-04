package ctx

import (
	"context"
	"net/http"

	"github.com/crypto-grill/app/internal/data"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"

	"github.com/crypto-grill/app/internal/config"
)

type ctxKey int

const (
	cfgCtxKey ctxKey = iota + 1

	usersCtxKey
	channelsCtxKey
	subscribersCtxKey
	subscribedChannelsCtxKey
	messagesCtxKey
	subscriptionProofsCtxKey
	hostCtxKey
	pubsubCtxKey
)

func SetConfig(cfg *config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, cfgCtxKey, cfg)
	}
}

func Config(r *http.Request) *config.Config {
	return r.Context().Value(cfgCtxKey).(*config.Config)
}

func SetHost(h host.Host) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, hostCtxKey, h)
	}
}

func Host(r *http.Request) host.Host {
	return r.Context().Value(hostCtxKey).(host.Host)
}

func SetPubSub(ps *pubsub.PubSub) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, pubsubCtxKey, ps)
	}
}

func PubSub(r *http.Request) *pubsub.PubSub {
	return r.Context().Value(pubsubCtxKey).(*pubsub.PubSub)
}

func SetUsers(d data.Users) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersCtxKey, d)
	}
}

func Users(r *http.Request) data.Users {
	return r.Context().Value(usersCtxKey).(data.Users)
}

func SetChannels(d data.Channels) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, channelsCtxKey, d)
	}
}

func Channels(r *http.Request) data.Channels {
	return r.Context().Value(channelsCtxKey).(data.Channels)
}

func SetSubscribers(d data.Subscribers) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, subscribersCtxKey, d)
	}
}

func Subscribers(r *http.Request) data.Subscribers {
	return r.Context().Value(subscribersCtxKey).(data.Subscribers)
}

func SetSubscribedChannels(d data.SubscribedChannels) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, subscribedChannelsCtxKey, d)
	}
}

func SubscribedChannels(r *http.Request) data.SubscribedChannels {
	return r.Context().Value(subscribedChannelsCtxKey).(data.SubscribedChannels)
}

func SetMessages(d data.Messages) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, messagesCtxKey, d)
	}
}

func Messages(r *http.Request) data.Messages {
	return r.Context().Value(messagesCtxKey).(data.Messages)
}

func SetSubscriptionProofs(d data.SubscriptionProofs) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, subscriptionProofsCtxKey, d)
	}
}

func SubscriptionProofs(r *http.Request) data.SubscriptionProofs {
	return r.Context().Value(subscriptionProofsCtxKey).(data.SubscriptionProofs)
}
