package storage

import (
	"errors"
	"time"

	"github.com/NordGus/shrtnr/domain/url/entities"

	"github.com/jmoiron/sqlx"
)

type record struct {
	ID        string `db:"id"`
	UUID      string `db:"uuid"`
	Target    string `db:"target"`
	CreatedAt int64  `db:"created_at"`
	DeletedAt int64  `db:"deleted_at"`
}

type Repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetByUUID(uuid entities.UUID) (entities.URL, error) {
	var (
		rcrd   record
		entity entities.URL
		term   = uuid.String()
	)

	err := repo.db.Get(&rcrd, "SELECT * FROM urls WHERE uuid = '?';", term)
	if err != nil {
		return entity, err
	}

	entity, err = entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0), time.Unix(rcrd.DeletedAt, 0))
	if err != nil {
		return entity, err
	}

	return entity, nil
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
