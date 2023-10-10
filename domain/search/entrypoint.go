package search

import (
	"context"
	"github.com/NordGus/shrtnr/domain/messagebus/url/created"
	"github.com/NordGus/shrtnr/domain/messagebus/url/deleted"
	"github.com/NordGus/shrtnr/domain/shared/trie"
	"github.com/NordGus/shrtnr/domain/storage"
	"sync"
)

var (
	cache      trie.Trie
	longsLimit int
	lock       sync.RWMutex
	ctx        context.Context

	repository Repository
)

func Start(parentCtx context.Context, maxConcurrency uint, searchLimit int) {
	ctx = parentCtx
	longsLimit = searchLimit

	cache = trie.NewTrie(maxConcurrency)
	repository = storage.GetURLRepository()

	created.Subscribe(onUrlCreatedSubscriber)
	deleted.Subscribe(onUrlDeletedSubscriber)
}
