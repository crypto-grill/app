package data

type Users interface {
	New() Users

	Transaction(func() error) error
}

type Channels interface {
	New() Channels

	Save(Channel) error

	Transaction(func() error) error
}

type Subscribers interface {
	New() Subscribers

	Transaction(func() error) error
}

type SubscribedChannels interface {
	New() SubscribedChannels

	Transaction(func() error) error
}

type Messages interface {
	New() Messages

	Save(Message) error

	Transaction(func() error) error
}

type SubscriptionProofs interface {
	New() SubscriptionProofs

	Transaction(func() error) error
}
