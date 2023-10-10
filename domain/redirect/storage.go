package redirect

import "github.com/NordGus/shrtnr/domain/storage/url"

type Repository interface {
	GetByShort(short string) (url.URL, error)
}
