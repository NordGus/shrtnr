package search

import (
	"context"
	"log"
	"strings"
	"sync"

	"github.com/NordGus/shrtnr/domain/shared/trie"
	"github.com/NordGus/shrtnr/domain/url/messagebus/created"
	"github.com/NordGus/shrtnr/domain/url/messagebus/deleted"
	"github.com/NordGus/shrtnr/domain/url/storage"
)

var (
	clearTargetCache trie.Trie
	fullTargetCache  trie.Trie
	shortCache       trie.Trie

	longsLimit int
	lock       sync.RWMutex
	ctx        context.Context

	repository  Repository
	redirectURL string
)

func Start(parentCtx context.Context, maxConcurrency uint, searchLimit int, redirectHost string, maxRecords uint) {
	ctx = parentCtx
	longsLimit = searchLimit
	redirectURL = redirectHost

	clearTargetCache = trie.NewTrie(maxConcurrency)
	fullTargetCache = trie.NewTrie(maxConcurrency)
	shortCache = trie.NewTrie(maxConcurrency)
	repository = storage.GetRepository()

	fillCaches(maxRecords)

	created.Subscribe(onUrlCreatedSubscriber)
	deleted.Subscribe(onUrlDeletedSubscriber)
}

func fillCaches(recordsLimit uint) {
	rcrds, err := repository.GetAllRecords(recordsLimit)
	if err != nil {
		log.Println(err) // ignores errors because there can be no records and still return an error
	}

	for _, rcrd := range rcrds {
		clearTargetCache.AddEntry(clearTargetEntry(rcrd.Target.String()))
		fullTargetCache.AddEntry(rcrd.Target.String())
		shortCache.AddEntry(rcrd.UUID.String())
	}
}

func clearTargetEntry(target string) string {
	clearEntry := strings.TrimPrefix(target, "https://")
	clearEntry = strings.TrimPrefix(clearEntry, "http://")
	clearEntry = strings.TrimPrefix(clearEntry, "www.")

	return clearEntry
}
