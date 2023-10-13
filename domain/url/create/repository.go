package create

import (
	"github.com/NordGus/shrtnr/domain/url"
)

type Repository interface {
	GetByUUID(uuid url.UUID) (url.URL, error)
	GetByTarget(target url.Target) (url.URL, error)
	CreateURL(entity url.URL) (url.URL, error)
	DeleteURL(id url.ID) (url.URL, error)
}
