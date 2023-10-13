package search

import (
	"github.com/NordGus/shrtnr/domain/url"
)

type Repository interface {
	GetURLsThatMatchTargets(matchTargets ...string) ([]url.URL, error)
}
