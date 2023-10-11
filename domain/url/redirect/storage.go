package redirect

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

type Repository interface {
	GetByShort(short string) (url.URL, error)
}
