package postgres

import (
	"context"
	"github.com/crypto-grill/app/internal/data"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
)

const subscriptionProofsTable = "subscription_proof"

type subscriptionProofs struct {
	db            *sqlx.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

func NewSubscriptionProofs(db *sqlx.DB) data.SubscriptionProofs {
	return &subscriptionProofs{
		db:            db,
		selectBuilder: sq.Select("*").From(subscriptionProofsTable).RunWith(db).PlaceholderFormat(sq.Dollar),
		deleteBuilder: sq.Delete(subscriptionProofsTable).RunWith(db).PlaceholderFormat(sq.Dollar),
	}
}

func (q *subscriptionProofs) New() data.SubscriptionProofs {
	return NewSubscriptionProofs(q.db)
}

func (q *subscriptionProofs) Transaction(fn func() error) error {
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
