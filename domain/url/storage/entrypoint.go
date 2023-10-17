package storage

import (
	"github.com/jmoiron/sqlx"
)

var (
	repository *Repository
)

func Start(db *sqlx.DB) error {
	repository = newRepository(db)

	return nil
}

func GetRepository() *Repository {
	return repository
}
