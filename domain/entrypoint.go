package domain

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NordGus/shrtnr/domain/url"
)

var (
	InitializationErr = errors.New("domain: failed to initialize")
)

func Start(ctx context.Context, env string, db *sql.DB, maxUrl uint, maxConcurrency uint, searchLimit int, redirectHost string) error {
	err := url.Start(ctx, env, db, maxUrl, maxConcurrency, searchLimit, redirectHost)
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	return nil
}
