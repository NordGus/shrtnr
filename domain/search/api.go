package search

import (
	"github.com/NordGus/shrtnr/domain/shared/response"
	"github.com/NordGus/shrtnr/domain/storage/url"
)

func SearchURLs(term string) ([]url.URL, error) {
	select {
	case <-ctx.Done():
		return []url.URL{}, ctx.Err()
	default:
		lock.RLock()
		defer lock.RUnlock()

		resp := buildSignal(term)
		resp = response.AndThen(resp, getLongsFromCache)
		resp = response.AndThen(resp, getRecordsFromRepository)

		return resp.urls, resp.err
	}
}
