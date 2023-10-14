package storage

import "database/sql"

var (
	repository *Repository
)

func Start(db *sql.DB) error {
	repository = newStorage(db)

	return nil
}

func GetRepository() *Repository {
	return repository
}
