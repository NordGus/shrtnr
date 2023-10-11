package search

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

type Repository interface {
	GetLikeLongs(linkLongs ...string) ([]url.URL, error)
}
