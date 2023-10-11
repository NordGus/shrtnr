package search

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
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
