package find

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetByID(id entities.ID) (entities.URL, error)
	GetByUUID(uuid entities.UUID) (entities.URL, error)
	GetAllInPage(page uint, perPage uint) ([]entities.URL, error)
}
