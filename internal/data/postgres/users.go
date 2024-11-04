package postgres

import (
	"context"
	"fmt"

	"github.com/crypto-grill/app/internal/data"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
)

const usersTable = "user_"

type users struct {
	db            *sqlx.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewUsers(db *sqlx.DB) data.Users {
	return &users{
		db:            db,
		selectBuilder: sq.Select("*").From(usersTable).RunWith(db).PlaceholderFormat(sq.Dollar),
		deleteBuilder: sq.Delete(usersTable).RunWith(db).PlaceholderFormat(sq.Dollar),
	}
}

func (q *users) New() data.Users {
	return NewUsers(q.db)
}

func (q *users) Save(user data.User) error {
	clauses := map[string]interface{}{
		"id":      user.ID,
		"pub_key": user.PubKey,
		"ip":      user.IP,
	}
	result := new(data.User)

	stmt := sq.Insert(usersTable).SetMap(clauses).RunWith(q.db).Suffix("RETURNING *").PlaceholderFormat(sq.Dollar)
	query, args, err := stmt.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build SQL query")
	}

	err = q.db.Get(result, query, args...)
	return errors.Wrap(err, "failed to execute insert query")
}

func (q *users) GetPubKeyForChannel(channelID int64) (string, error) {
	var pubKey string
	queryBuilder := q.selectBuilder.
		Columns("u.pub_key").
		From("user_ u").
		Join("channel c ON u.id = c.sender_id").
		Where(sq.Eq{"c.id": channelID})

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build SQL query: %w", err)
	}

	err = q.db.Get(&pubKey, query, args...)
	if err != nil {
		return "", fmt.Errorf("failed to get public key for channel %d: %w", channelID, err)
	}
	return pubKey, nil
}

func (q *users) GetIPsForChannels(channelIDs []int64) ([]string, error) {
	var ips []string

	// Maybe we can do better
	query := `
		SELECT u.ip
		FROM user_ u
		JOIN channel c ON u.id = c.sender_id
		WHERE c.id = ANY($1)
		ORDER BY array_position($1, c.id)
	`

	err := q.db.Select(&ips, query, pq.Array(channelIDs))
	if err != nil {
		return nil, fmt.Errorf("failed to get IPs for channels %v: %w", channelIDs, err)
	}
	return ips, nil
}

func (q *users) GetIPsForSubsriber(subscriberIDs []int64) ([]string, error) {
	var ips []string

	// Maybe we can do better
	query := `
		SELECT u.ip
		FROM user_ u
		JOIN subscriber s ON u.id = s.user_id
		WHERE s.id = ANY($1)
		ORDER BY array_position($1, s.id)
	`

	err := q.db.Select(&ips, query, pq.Array(subscriberIDs))
	if err != nil {
		return nil, fmt.Errorf("failed to get IPs for channels %v: %w", subscriberIDs, err)
	}
	return ips, nil
}

func (q *users) Transaction(fn func() error) error {
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
