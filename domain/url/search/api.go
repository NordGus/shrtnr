package search

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

func ByTerm(term string) ([]entities.URL, error) {
	select {
	case <-ctx.Done():
		return []entities.URL{}, ctx.Err()
	default:
		resp := buildSearchURLsResponse(term)
		resp = railway.AndThen(resp, getMatchersFromClearTargetCache)
		resp = railway.AndThen(resp, getMatchersFromFullTargetCache)
		resp = railway.AndThen(resp, getMatchersFromShortCache)
		resp = railway.AndThen(resp, getRecordsFromRepository)

		return resp.records, resp.err
	}
}
