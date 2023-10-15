package storage

import (
	"errors"

	"github.com/NordGus/shrtnr/domain/url/entities"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func newStorage(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetByUUID(uuid entities.UUID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) GetByTarget(target entities.Target) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) CreateURL(entity entities.URL) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) DeleteURL(id entities.ID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) GetURLsThatMatchTargets(matchTargets ...string) ([]entities.URL, error) {
	// TODO: Implement
	return []entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) GetByID(id entities.ID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) GetAllInPage(page uint, perPage uint) ([]entities.URL, error) {
	// TODO: Implement
	return []entities.URL{}, errors.New("storage: not implemented")
}
