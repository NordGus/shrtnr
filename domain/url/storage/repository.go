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
}

type Repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetAllRecords returns requested records in the database ordered from newest to oldest.
//
// When limit is 0 (zero) it returns all records in the database.
func (repo *Repository) GetAllRecords(limit uint) ([]entities.URL, error) {
	var (
		err   error
		rcrds = make([]record, 0)
		ents  = make([]entities.URL, 0)
	)

	err = repo.db.Select(&rcrds, "SELECT * FROM urls ORDER BY created_at DESC LIMIT ?", limit)
	if err != nil {
		return ents, err
	}

	ents = make([]entities.URL, 0, len(rcrds))

	for _, rcrd := range rcrds {
		u, err := entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0))
		if err != nil {
			return ents, err
		}

		ents = append(ents, u)
	}

	return ents, nil
}

func (repo *Repository) GetByID(id entities.ID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) GetByUUID(uuid entities.UUID) (entities.URL, error) {
	var (
		rcrd   record
		entity entities.URL
		term   = uuid.String()
	)

	err := repo.db.Get(&rcrd, "SELECT * FROM urls WHERE uuid = ?;", term)
	if err != nil {
		return entity, err
	}

	entity, err = entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0))
	if err != nil {
		return entities.URL{}, err
	}

	return entity, nil
}

func (repo *Repository) GetByTarget(target entities.Target) (entities.URL, error) {
	var (
		rcrd   record
		term   = target.String()
		entity entities.URL
	)

	err := repo.db.Get(&rcrd, "SELECT * FROM urls WHERE target = ?", term)
	if err != nil {
		return entity, err
	}

	entity, err = entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0))
	if err != nil {
		return entities.URL{}, err
	}

	return entity, nil
}

func (repo *Repository) CreateURL(entity entities.URL) (entities.URL, error) {
	var (
		id        = entity.ID.String()
		uuid      = entity.UUID.String()
		target    = entity.Target.String()
		createdAt = entity.CreatedAt.Unix()
	)

	_, err := repo.db.Exec("INSERT INTO urls (id, uuid, target, created_at) VALUES (?, ?, ?, ?)", id, uuid, target, createdAt)
	if err != nil {
		return entities.URL{}, err
	}

	entity, err = entities.NewURL(id, uuid, target, time.Unix(createdAt, 0))
	if err != nil {
		return entities.URL{}, err
	}

	return entity, nil
}

func (repo *Repository) DeleteURL(id entities.ID) (entities.URL, error) {
	// TODO: Implement
	return entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) GetURLsThatMatchTargets(matchTargets ...string) ([]entities.URL, error) {
	// TODO: Implement
	return []entities.URL{}, errors.New("storage: not implemented")
}

func (repo *Repository) GetAllInPage(page uint, perPage uint) ([]entities.URL, error) {
	var (
		err   error
		rcrds = make([]record, 0, perPage)
		ents  = make([]entities.URL, 0, perPage)
	)

	err = repo.db.Select(&rcrds, "SELECT * FROM urls ORDER BY created_at DESC LIMIT ? OFFSET ?", perPage, (page-1)*perPage)
	if err != nil {
		return ents, err
	}

	for _, rcrd := range rcrds {
		u, err := entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0))
		if err != nil {
			return ents, err
		}

		ents = append(ents, u)
	}

	return ents, nil
}
