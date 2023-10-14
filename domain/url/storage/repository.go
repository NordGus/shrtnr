package storage

import (
	"database/sql"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository struct {
	db *sql.DB
}

func newStorage(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (db *Repository) GetByUUID(uuid entities.UUID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, nil
}

func (db *Repository) GetByTarget(target entities.Target) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, nil
}

func (db *Repository) CreateURL(entity entities.URL) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, nil
}

func (db *Repository) DeleteURL(id entities.ID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, nil
}

func (db *Repository) GetURLsThatMatchTargets(matchTargets ...string) ([]entities.URL, error) {
	// TODO: Implement
	return []entities.URL{}, nil
}

func (db *Repository) GetByID(id entities.ID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, nil
}

func (db *Repository) GetAllInPage(page uint, perPage uint) ([]entities.URL, error) {
	// TODO: Implement
	return []entities.URL{}, nil
}
