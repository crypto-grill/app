package data

import "time"

type User struct {
	ID     int64  `db:"id"`
	PubKey string `db:"pub_key"`
	IP     string `db:"ip"`
}

type Channel struct {
	ID        int64      `db:"id"`
	SenderID  int64      `db:"sender_id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
}

type Subscriber struct {
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
	ExpiresAt *time.Time `db:"expires_at"`
}
