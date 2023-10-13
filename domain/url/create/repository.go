package create

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetByUUID(uuid entities.UUID) (entities.URL, error)
	GetByTarget(target entities.Target) (entities.URL, error)
	CreateURL(entity entities.URL) (entities.URL, error)
	DeleteURL(id entities.ID) (entities.URL, error)
}
