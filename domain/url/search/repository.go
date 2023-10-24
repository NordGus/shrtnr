package search

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetAllRecords(limit uint) ([]entities.URL, error)
	GetURLsThatMatchTargets(matchTargets ...string) ([]entities.URL, error)
}
