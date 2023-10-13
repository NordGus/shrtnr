package redirect

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetByUUID(uuid entities.UUID) (entities.URL, error)
}
