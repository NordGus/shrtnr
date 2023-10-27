package search

import (
	"errors"
	"fmt"
	"github.com/NordGus/shrtnr/domain/shared/trie"
	"github.com/NordGus/shrtnr/domain/url/entities"
	"log"
	"strings"
	"sync"
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
	return searchURLsResponse{
		term:     term,
		matchers: make([]string, 0, longsLimit*4),
		records:  make([]entities.URL, 0, longsLimit),
	}
}

func getMatchersFromClearTargetCache(response searchURLsResponse) searchURLsResponse {
	term := strings.ToLower(response.term)
	term = strings.TrimPrefix(term, "https://")
	term = strings.TrimPrefix(term, "http://")

	matchers, err := clearTargetCache.FindEntries(term, longsLimit)
	if err != nil && !errors.Is(err, trie.EntryNotPresentErr) {
		response.err = errors.Join(response.err, err)
	}

	for _, matcher := range matchers {
		response.matchers = append(response.matchers, "%"+matcher)
	}

	return response
}

func getMatchersFromFullTargetCache(response searchURLsResponse) searchURLsResponse {
	term := strings.ToLower(response.term)

	matchers, err := fullTargetCache.FindEntries(term, longsLimit)
	if err != nil && !errors.Is(err, trie.EntryNotPresentErr) {
		response.err = errors.Join(response.err, err)
	}

	response.matchers = append(response.matchers, matchers...)

	return response
}

func getMatchersFromShortCache(response searchURLsResponse) searchURLsResponse {
	term := strings.TrimPrefix(response.term, fmt.Sprintf("%s/r/", redirectURL))

	matchers, err := shortCache.FindEntries(term, longsLimit)
	if err != nil && !errors.Is(err, trie.EntryNotPresentErr) {
		response.err = errors.Join(response.err, err)
	}

	response.matchers = append(response.matchers, matchers...)

	return response
}

func getRecordsFromRepository(response searchURLsResponse) searchURLsResponse {
	var (
		count int

		resultsCh = make(chan []entities.URL, 3)
		wg        = new(sync.WaitGroup)
		added     = make(map[entities.ID]bool)
	)

	wg.Add(3)

	go func(wg *sync.WaitGroup, resultsCh chan<- []entities.URL, matchers []string) {
		defer wg.Done()

		records, err := repository.GetURLsLikeTargets(uint(longsLimit), matchers...)
		if err != nil {
			log.Println(err)
		}

		resultsCh <- records
	}(wg, resultsCh, response.matchers)

	go func(wg *sync.WaitGroup, resultsCh chan<- []entities.URL, matchers []string) {
		defer wg.Done()

		records, err := repository.GetURLsByTargets(uint(longsLimit), matchers...)
		if err != nil {
			log.Println(err)
		}

		resultsCh <- records
	}(wg, resultsCh, response.matchers)

	go func(wg *sync.WaitGroup, resultsCh chan<- []entities.URL, matchers []string) {
		defer wg.Done()

		records, err := repository.GetURLsByUUIDs(uint(longsLimit), matchers...)
		if err != nil {
			log.Println(err)
		}

		resultsCh <- records
	}(wg, resultsCh, response.matchers)

	wg.Wait()
	close(resultsCh)

	for results := range resultsCh {
		for _, result := range results {
			if count == longsLimit {
				break
			}

			if _, ok := added[result.ID]; ok {
				continue
			}

			added[result.ID] = true
			response.records = append(response.records, result)
			count++
		}
	}

	return response
}
