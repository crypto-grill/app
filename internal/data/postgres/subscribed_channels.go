package postgres

import (
	"context"

	"github.com/crypto-grill/app/internal/data"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
)

const subscribedChannelsTable = "subscribed_channel"

type subscribedChannels struct {
	db            *sqlx.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewSubscribedChannels(db *sqlx.DB) data.SubscribedChannels {
	return &subscribedChannels{
		db:            db,
		selectBuilder: sq.Select("*").From(subscribedChannelsTable).RunWith(db).PlaceholderFormat(sq.Dollar),
		deleteBuilder: sq.Delete(subscribedChannelsTable).RunWith(db).PlaceholderFormat(sq.Dollar),
	}
}

func (q *subscribedChannels) New() data.SubscribedChannels {
	return NewSubscribedChannels(q.db)
}

func (q *subscribedChannels) Transaction(fn func() error) error {
	tx, err := q.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}

	if err := fn(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Wrapf(rbErr, "transaction rollback failed after error: %v", err)
		}
		return errors.Wrap(err, "transaction failed")
	}

	err = tx.Commit()

	return errors.Wrap(err, "failed to commit transaction")
}
