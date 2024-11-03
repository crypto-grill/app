package postgres

import (
	"context"

	"github.com/crypto-grill/app/internal/data"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
)

const channelsTable = "channel"

type channels struct {
	db            *sqlx.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewChannels(db *sqlx.DB) data.Channels {
	return &channels{
		db:            db,
		selectBuilder: sq.Select("*").From(channelsTable).RunWith(db).PlaceholderFormat(sq.Dollar),
		deleteBuilder: sq.Delete(channelsTable).RunWith(db).PlaceholderFormat(sq.Dollar),
	}
}

func (q *channels) New() data.Channels {
	return NewChannels(q.db)
}

func (q *channels) Save(msg data.Channel) error {
	clauses := map[string]interface{}{
		"id":        msg.ID,
		"sender_id": msg.SenderID,
		"name":      msg.Name,
	}

	result := new(data.Channel)

	stmt := sq.Insert(channelsTable).SetMap(clauses).RunWith(q.db).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := stmt.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build SQL query")
	}

	err = q.db.Get(result, query, args...)

	return errors.Wrap(err, "failed to execute insert query")
}

func (q *channels) Transaction(fn func() error) error {
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
