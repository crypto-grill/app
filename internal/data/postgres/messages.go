package postgres

import (
	"context"
	"time"

	"github.com/crypto-grill/app/internal/data"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
)

const messagesTable = "message"

type messages struct {
	db            *sqlx.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewMessages(db *sqlx.DB) data.Messages {
	return &messages{
		db:            db,
		selectBuilder: sq.Select("*").From(messagesTable).RunWith(db).PlaceholderFormat(sq.Dollar),
		deleteBuilder: sq.Delete(messagesTable).RunWith(db).PlaceholderFormat(sq.Dollar),
	}
}

func (q *messages) New() data.Messages {
	return NewMessages(q.db)
}

func (q *messages) InChannel(channelID int64) data.Messages {
	newSelectBuilder := q.selectBuilder.Where("channel_id = ?", channelID)
	return &messages{
		db:            q.db,
		selectBuilder: newSelectBuilder,
		deleteBuilder: q.deleteBuilder,
	}
}

func (q *messages) After(time time.Time) data.Messages {
	newSelectBuilder := q.selectBuilder.Where("created_at > ?", time)
	return &messages{
		db:            q.db,
		selectBuilder: newSelectBuilder,
		deleteBuilder: q.deleteBuilder,
	}
}

func (q *messages) Save(msg data.Message) error {
	clauses := map[string]interface{}{
		"id":         msg.ID,
		"channel_id": msg.ChannelID,
		"message":    msg.Message,
		"created_at": msg.CreatedAt,
	}
	result := new(data.Message)

	stmt := sq.Insert(messagesTable).SetMap(clauses).RunWith(q.db).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := stmt.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build SQL query")
	}

	err = q.db.Get(result, query, args...)

	return errors.Wrap(err, "failed to execute insert query")
}

func (q *messages) Select() ([]data.Message, error) {
	stmt := sq.Select("*").From(messagesTable).RunWith(q.db).PlaceholderFormat(sq.Dollar)
	query, args, err := stmt.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build SQL query")
	}

	var messages []data.Message
	err = q.db.Select(&messages, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute select query")
	}
	return messages, nil
}

func (q *messages) Transaction(fn func() error) error {
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
