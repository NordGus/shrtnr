package create

import "github.com/NordGus/shrtnr/server/storage/url"

type Repository interface {
	GetByShort(short string) (url.URL, error)
	GetByFull(full string) (url.URL, error)
	CreateURL(short string, full string) (url.URL, error)
	DeleteURL(short string) (url.URL, error)
}