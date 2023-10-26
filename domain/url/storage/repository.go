package storage

import (
	"strings"
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
func (repo *Repository) GetAllRecords(limit uint) ([]entities.URL, error) {
	var (
		rcrds = make([]record, 0, limit)
		ents  = make([]entities.URL, 0, limit)
	)

	err := repo.db.Select(&rcrds, "SELECT * FROM urls ORDER BY created_at DESC LIMIT ?", limit)
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

func (repo *Repository) GetByID(id entities.ID) (entities.URL, error) {
	var (
		rcrd = record{ID: id.String()}

		entity entities.URL
	)

	err := repo.db.Get(&rcrd, "SELECT * FROM urls WHERE id = ?;", rcrd.ID)
	if err != nil {
		return entity, err
	}

	entity, err = entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0))
	if err != nil {
		return entities.URL{}, err
	}

	return entity, nil
}

func (repo *Repository) GetByUUID(uuid entities.UUID) (entities.URL, error) {
	var (
		rcrd = record{UUID: uuid.String()}

		entity entities.URL
	)

	err := repo.db.Get(&rcrd, "SELECT * FROM urls WHERE uuid = ?;", rcrd.UUID)
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
		rcrd = record{Target: target.String()}

		entity entities.URL
	)

	err := repo.db.Get(&rcrd, "SELECT * FROM urls WHERE target = ?", rcrd.Target)
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
	var rcrd = record{
		ID:        entity.ID.String(),
		UUID:      entity.UUID.String(),
		Target:    entity.Target.String(),
		CreatedAt: entity.CreatedAt.Unix(),
	}

	_, err := repo.db.NamedExec("INSERT INTO urls (id, uuid, target, created_at) VALUES (:id, :uuid, :target, :created_at)", rcrd)
	if err != nil {
		return entities.URL{}, err
	}

	entity, err = entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0))
	if err != nil {
		return entities.URL{}, err
	}

	return entity, nil
}

func (repo *Repository) DeleteURL(id entities.ID) (entities.URL, error) {
	var rcrd = record{ID: id.String()}

	err := repo.db.Get(&rcrd, "SELECT * FROM urls WHERE id = ?", rcrd.ID)
	if err != nil {
		return entities.URL{}, err
	}

	_, err = repo.db.NamedExec("DELETE FROM urls WHERE id = :id", rcrd)
	if err != nil {
		return entities.URL{}, err
	}

	entity, err := entities.NewURL(rcrd.ID, rcrd.UUID, rcrd.Target, time.Unix(rcrd.CreatedAt, 0))
	if err != nil {
		return entities.URL{}, err
	}

	return entity, nil
}

func (repo *Repository) GetURLsLikeTargets(limit uint, targets ...string) ([]entities.URL, error) {
	if len(targets) == 0 {
		return []entities.URL{}, nil
	}

	var (
		results      []entities.URL
		queryBuilder strings.Builder

		rcrds  = make([]record, 0, limit)
		params = make([]interface{}, len(targets), len(targets)+1)
	)

	queryBuilder.WriteString("SELECT * FROM urls WHERE ")

	for i, target := range targets {
		queryBuilder.WriteString("target LIKE ? ")
		params[i] = target

		if i < len(targets)-1 {
			queryBuilder.WriteString("OR ")
		}
	}

	queryBuilder.WriteString("ORDER BY created_at DESC LIMIT ?")
	params = append(params, limit)

	err := repo.db.Select(&rcrds, queryBuilder.String(), params...)
	if err != nil {
		return []entities.URL{}, err
	}

	results = make([]entities.URL, len(rcrds))

	for i, r := range rcrds {
		u, err := entities.NewURL(r.ID, r.UUID, r.Target, time.Unix(r.CreatedAt, 0))
		if err != nil {
			return []entities.URL{}, err
		}

		results[i] = u
	}

	return results, nil
}

func (repo *Repository) GetURLsByTargets(limit uint, targets ...string) ([]entities.URL, error) {
	if len(targets) == 0 {
		return []entities.URL{}, nil
	}

	var (
		rcrds = make([]record, 0, limit)

		results []entities.URL
	)

	query, args, err := sqlx.In("SELECT * FROM urls WHERE target IN(?) ORDER BY created_at DESC LIMIT ?", targets, limit)
	if err != nil {
		return results, err
	}

	err = repo.db.Select(&rcrds, query, args...)
	if err != nil {
		return results, err
	}

	results = make([]entities.URL, len(rcrds))

	for i, r := range rcrds {
		u, err := entities.NewURL(r.ID, r.UUID, r.Target, time.Unix(r.CreatedAt, 0))
		if err != nil {
			return []entities.URL{}, err
		}

		results[i] = u
	}

	return results, nil
}

func (repo *Repository) GetURLsByUUIDs(limit uint, uuids ...string) ([]entities.URL, error) {
	if len(uuids) == 0 {
		return []entities.URL{}, nil
	}

	var (
		rcrds   = make([]record, 0, limit)
		results []entities.URL
	)

	query, args, err := sqlx.In("SELECT * FROM urls WHERE uuid IN(?) ORDER BY created_at DESC LIMIT ?", uuids, limit)
	if err != nil {
		return results, err
	}

	err = repo.db.Select(&rcrds, query, args...)
	if err != nil {
		return results, err
	}

	results = make([]entities.URL, len(rcrds))

	for i, r := range rcrds {
		u, err := entities.NewURL(r.ID, r.UUID, r.Target, time.Unix(r.CreatedAt, 0))
		if err != nil {
			return []entities.URL{}, err
		}

		results[i] = u
	}

	return results, nil
}

func (repo *Repository) GetAllInPage(page uint, perPage uint) ([]entities.URL, error) {
	var (
		rcrds = make([]record, 0, perPage)
		ents  = make([]entities.URL, 0, perPage)

		limit  = perPage
		offset = (page - 1) * perPage
	)

	err := repo.db.Select(&rcrds, "SELECT * FROM urls ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset)
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
