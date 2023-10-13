package search

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type Repository interface {
	GetURLsThatMatchTargets(matchTargets ...string) ([]entities.URL, error)
}
