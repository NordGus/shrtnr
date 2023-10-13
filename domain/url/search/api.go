package search

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

func SearchURLs(term string) ([]url.URL, error) {
	select {
	case <-ctx.Done():
		return []url.URL{}, ctx.Err()
	default:
		lock.RLock()
		defer lock.RUnlock()

		resp := buildSignal(term)
		resp = railway.AndThen(resp, getLongsFromCache)
		resp = railway.AndThen(resp, getRecordsFromRepository)

		return resp.urls, resp.err
	}
}
