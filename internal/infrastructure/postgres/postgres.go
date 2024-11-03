package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"time"
)

func New(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to postgres")
	}

	db.SetConnMaxLifetime(time.Minute)

	return db, nil
}
