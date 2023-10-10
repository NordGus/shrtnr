package search

import "github.com/NordGus/shrtnr/domain/storage/url"

type Repository interface {
	GetLikeLongs(linkLongs ...string) ([]url.URL, error)
}
