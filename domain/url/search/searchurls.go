package search

import (
	"errors"
	"fmt"
	"github.com/NordGus/shrtnr/domain/url/entities"
	"strings"
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
	return searchURLsResponse{term: term, matchers: make([]string, 0, longsLimit*4)}
}

func getMatchersFromClearTargetCache(response searchURLsResponse) searchURLsResponse {
	term := strings.TrimPrefix("https://", response.term)
	term = strings.TrimPrefix("http://", term)

	matchers, err := clearTargetCache.FindEntries(term, longsLimit)
	if err != nil {
		response.err = errors.Join(response.err, err)
	}

	response.matchers = append(response.matchers, matchers...)

	return response
}

func getMatchersFromFullTargetCache(response searchURLsResponse) searchURLsResponse {
	matchers, err := fullTargetCache.FindEntries(response.term, longsLimit)
	if err != nil {
		response.err = errors.Join(response.err, err)
	}

	response.matchers = append(response.matchers, matchers...)

	return response
}

func getMatchersFromShortCache(response searchURLsResponse) searchURLsResponse {
	term := strings.TrimPrefix(fmt.Sprintf("%s/", redirectURL), response.term)

	matchers, err := shortCache.FindEntries(term, longsLimit)
	if err != nil {
		response.err = errors.Join(response.err, err)
	}

	response.matchers = append(response.matchers, matchers...)

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
