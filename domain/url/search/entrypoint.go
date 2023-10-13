package search

import (
	"context"
	"sync"

	"github.com/NordGus/shrtnr/domain/shared/trie"
	"github.com/NordGus/shrtnr/domain/url/messagebus/created"
	"github.com/NordGus/shrtnr/domain/url/messagebus/deleted"
	"github.com/NordGus/shrtnr/domain/url/storage"
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
	repository = storage.GetRepository()

	created.Subscribe(onUrlCreatedSubscriber)
	deleted.Subscribe(onUrlDeletedSubscriber)
}
