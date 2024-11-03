package postgres

import (
	"context"

	"github.com/crypto-grill/app/internal/data"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
)

const subscribersTable = "subscriber"

type subscribers struct {
	db            *sqlx.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewSubscribers(db *sqlx.DB) data.Subscribers {
	return &subscribers{
		db:            db,
		selectBuilder: sq.Select("*").From(subscribersTable).RunWith(db).PlaceholderFormat(sq.Dollar),
		deleteBuilder: sq.Delete(subscribersTable).RunWith(db).PlaceholderFormat(sq.Dollar),
	}
}

func (q *subscribers) New() data.Subscribers {
	return NewSubscribers(q.db)
}

func (q *subscribers) Save(subscriber data.Subscriber) error {
	clauses := map[string]interface{}{
		"user_id":    subscriber.UserID,
		"channel_id": subscriber.ChannelID,
	}
	result := new(data.Subscriber)

	stmt := sq.Insert(subscribersTable).SetMap(clauses).RunWith(q.db).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := stmt.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build SQL query")
	}

	err = q.db.Get(result, query, args...)

	return errors.Wrap(err, "failed to execute insert query")
}

func (q *subscribers) Transaction(fn func() error) error {
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
