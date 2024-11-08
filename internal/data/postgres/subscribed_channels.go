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

func (q *subscribedChannels) Save(msg data.SubscribedChannel) error {
	clauses := map[string]interface{}{
		"channel_id": msg.ChannelID,
	}
	result := new(data.Message)

	stmt := sq.Insert(subscribedChannelsTable).SetMap(clauses).RunWith(q.db).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := stmt.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build SQL query")
	}

	err = q.db.Get(result, query, args...)

	return errors.Wrap(err, "failed to execute insert query")
}

func (q *subscribedChannels) SelectChannelIDs() ([]int64, error) {
	queryBuilder := q.selectBuilder.Columns("channel_id")
	var channelIDs []int64

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build SQL query")
	}

	err = q.db.Select(&channelIDs, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute select query")
	}
	return channelIDs, nil
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
