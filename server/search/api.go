package search

import (
	"github.com/NordGus/shrtnr/server/shared/response"
	"github.com/NordGus/shrtnr/server/storage/url"
)

type signal struct {
	term  string
	longs []string
	urls  []url.URL
	err   error
}

func (s signal) Error() error {
	return s.err
}

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

func buildSignal(term string) signal {
	return signal{term: term}
}

func getLongsFromCache(sig signal) signal {
	sig.longs, sig.err = cache.FindEntries(sig.term, longsLimit)

	return sig
}

func getRecordsFromRepository(sig signal) signal {
	sig.urls, sig.err = repository.GetLikeLongs(sig.longs...)

	return sig
}
