package search

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

func SearchURLs(term string) ([]entities.URL, error) {
	select {
	case <-ctx.Done():
		return []entities.URL{}, ctx.Err()
	default:
		lock.RLock()
		defer lock.RUnlock()

		resp := buildSearchURLsResponse(term)
		resp = railway.AndThen(resp, getMatchersFromCache)
		resp = railway.AndThen(resp, getRecordsFromRepository)

		return resp.records, resp.err
	}
}
