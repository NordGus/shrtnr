package redirect

import (
	"github.com/NordGus/shrtnr/domain/url"
)

type Repository interface {
	GetByUUID(uuid url.UUID) (url.URL, error)
}
