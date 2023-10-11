package storage

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

type URLRepository interface {
	GetByShort(short string) (url.URL, error)
	GetByFull(full string) (url.URL, error)
	CreateURL(short string, full string) (url.URL, error)
	DeleteURL(id uint) (url.URL, error)

	GetLikeLongs(linkLongs ...string) ([]url.URL, error)

	GetByID(id uint) (url.URL, error)
	GetAllInPage(page uint, perPage uint) ([]url.URL, error)
}
