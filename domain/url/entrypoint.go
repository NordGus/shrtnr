package url

import (
	"context"
	"errors"

	"github.com/NordGus/shrtnr/domain/url/create"
	"github.com/NordGus/shrtnr/domain/url/find"
	"github.com/NordGus/shrtnr/domain/url/messagebus"
	"github.com/NordGus/shrtnr/domain/url/search"
	"github.com/NordGus/shrtnr/domain/url/storage"

	"github.com/jmoiron/sqlx"
)

var (
	InitializationErr = errors.New("url: failed to initialize")
)

// Start initializes all services in the domain
func Start(ctx context.Context, env string, db *sqlx.DB, maxUrl uint, maxConcurrency uint, searchLimit int) error {
	messagebus.Start(ctx)

	err := storage.Start(db)
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	create.Start(ctx, maxUrl)
	search.Start(ctx, maxConcurrency, searchLimit)

	err = find.Start(ctx)
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	return nil
}
