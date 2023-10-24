package search

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetAllRecords(limit uint) ([]entities.URL, error)
	GetURLsLikeTargets(limit uint, targets ...string) ([]entities.URL, error)
	GetURLsByTargets(limit uint, targets ...string) ([]entities.URL, error)
	GetURLsByUUIDs(limit uint, uuids ...string) ([]entities.URL, error)
}
