package find

import (
	"github.com/NordGus/shrtnr/domain/url"
)

type Repository interface {
	GetByID(id string) (url.URL, error)
	GetAllInPage(page uint, perPage uint) ([]url.URL, error)
}
