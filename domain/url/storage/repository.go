package storage

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetByUUID(uuid entities.UUID) (entities.URL, error)
	GetByTarget(target entities.Target) (entities.URL, error)
	CreateURL(entity entities.URL) (entities.URL, error)
	DeleteURL(id entities.ID) (entities.URL, error)

	GetURLsThatMatchTargets(matchTargets ...string) ([]entities.URL, error)

	GetByID(id entities.ID) (entities.URL, error)
	GetAllInPage(page uint, perPage uint) ([]entities.URL, error)
}
