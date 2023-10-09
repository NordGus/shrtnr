package redirect

import "github.com/NordGus/shrtnr/server/storage/url"

type Repository interface {
	GetByShort(short string) (url.URL, error)
}
