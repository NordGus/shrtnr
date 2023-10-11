package find

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

type Repository interface {
	GetByID(id uint) (url.URL, error)
	GetAllInPage(page uint, perPage uint) ([]url.URL, error)
}
