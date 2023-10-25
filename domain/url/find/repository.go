package find

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetByUUID(uuid entities.UUID) (entities.URL, error)
	GetAllInPage(page uint, perPage uint) ([]entities.URL, error)
}
