package data

import "time"

type Users interface {
	New() Users

	Save(User) error
	GetPubKeyForChannel(int64) (string, error)
	GetIPsForChannels([]int64) ([]string, error) // get sender ips
	GetIPsForSubsriber([]int64) ([]string, error)

	Transaction(func() error) error
}

type Channels interface {
	New() Channels

	Save(Channel) error
	GetName(int64) (string, error)
	GetSender(channelID int64) (int64, error)
	Select() ([]Channel, error)

	Transaction(func() error) error
}

type Subscribers interface {
	New() Subscribers

	FilterByChannelID(int64) Subscribers

	Select() ([]Subscriber, error)
	Save(Subscriber) error

	Transaction(func() error) error
}

type SubscribedChannels interface {
	New() SubscribedChannels

	Save(SubscribedChannel) error
	SelectChannelIDs() ([]int64, error)

	Transaction(func() error) error
}

type Messages interface {
	New() Messages

	InChannel(int64) Messages
	After(time.Time) Messages

	Save(Message) error
	Select() ([]Message, error)

	Transaction(func() error) error
}

type SubscriptionProofs interface {
	New() SubscriptionProofs

	Unexpired() SubscriptionProofs
	InChannels([]int64) SubscriptionProofs

	Select() ([]SubscriptionProof, error)

	Transaction(func() error) error
}
