package data

import "time"

type User struct {
	ID     int64  `db:"id"`
	PubKey string `db:"pub_key"`
}

type Channel struct {
	ID       int64  `db:"id"`
	SenderID int64  `db:"sender_id"`
	Name     string `db:"name"`
}

type Subscriber struct {
	ID        int64 `db:"id"`
	UserID    int64 `db:"user_id"`
	ChannelID int64 `db:"channel_id"`
}

type SubscribedChannel struct {
	ID        int64 `db:"id"`
	ChannelID int64 `db:"channel_id"`
}

type Message struct {
	ID        int64      `db:"id"`
	ChannelID int64      `db:"channel_id"`
	Message   string     `db:"message"`
	CreatedAt *time.Time `db:"created_at"`
}

type SubscriptionProof struct {
	ID        int64      `db:"id"`
	ChannelID int64      `db:"channel_id"`
	Signature string     `db:"signature"`
	Message   string     `db:"message"`
	CreatedAt *time.Time `db:"created_at"`
}
