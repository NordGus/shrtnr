package remove

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetByID(id entities.ID) (entities.URL, error)
	DeleteURL(id entities.ID) (entities.URL, error)
	CreateURL(entity entities.URL) (entities.URL, error)
}
