package search

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type searchURLsResponse struct {
	term     string
	matchers []string
	records  []entities.URL
	err      error
}

func (s searchURLsResponse) Success() bool {
	return s.err == nil
}

func buildSearchURLsResponse(term string) searchURLsResponse {
	return searchURLsResponse{term: term}
}

func getMatchersFromCache(response searchURLsResponse) searchURLsResponse {
	matchers, err := cache.FindEntries(response.term, longsLimit)
	if err != nil {
		response.err = errors.Join(response.err, err)
	}

	response.matchers = matchers

	return response
}

func getRecordsFromRepository(response searchURLsResponse) searchURLsResponse {
	records, err := repository.GetURLsThatMatchTargets(response.matchers...)
	if err != nil {
		response.err = errors.Join(response.err, err)
	}

	response.records = records

	return response
}
