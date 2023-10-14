package domain

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"

	"github.com/NordGus/shrtnr/domain/url"
)

var (
	InitializationErr = errors.New("domain: failed to initialize")
)

func Start(ctx context.Context, env string, db *sqlx.DB, maxUrl uint, maxConcurrency uint, searchLimit int, redirectHost string) error {
	err := url.Start(ctx, env, db, maxUrl, maxConcurrency, searchLimit, redirectHost)
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	return nil
}
