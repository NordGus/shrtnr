package storage

import (
	"errors"
	"log"
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

	entity, err = entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0), time.Unix(rcrd.DeletedAt, 0))
	if err != nil {
		return entity, err
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

	entity, err = entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0), time.Unix(rcrd.DeletedAt, 0))
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (repo *Repository) CreateURL(entity entities.URL) (entities.URL, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return entities.URL{}, err
	}

	_, err = tx.Exec(
		"INSERT INTO urls (id, uuid, target, created_at, deleted_at) VALUES (?, ?, ?, ?, ?)", entity.ID.String(),
		entity.UUID.String(), entity.Target.String(), entity.CreatedAt.Unix(), entity.DeletedAt.Unix())
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalln(err)
		}

		return entities.URL{}, err
	}

	err = tx.Commit()
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
		u, err := entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0), time.Unix(rcrd.DeletedAt, 0))
		if err != nil {
			return ents, err
		}

		ents = append(ents, u)
	}

	return ents, nil
}
