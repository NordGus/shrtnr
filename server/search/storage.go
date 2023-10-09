package search

import "github.com/NordGus/shrtnr/server/storage/url"

type Repository interface {
	GetLikeLongs(linkLongs ...string) ([]url.URL, error)
}
