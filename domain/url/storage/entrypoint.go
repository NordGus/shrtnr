package storage

import (
	"github.com/jmoiron/sqlx"
)

var (
	repository *Repository
)

func Start(db *sqlx.DB) error {
	repository = newStorage(db)

	return nil
}

func GetRepository() *Repository {
	return repository
}
