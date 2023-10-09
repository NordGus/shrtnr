package search

import (
	"context"
	"errors"
	"github.com/NordGus/shrtnr/server/messagebus/url/created"
	"github.com/NordGus/shrtnr/server/messagebus/url/deleted"
	"github.com/NordGus/shrtnr/server/shared/trie"
	"github.com/NordGus/shrtnr/server/storage"
	"github.com/NordGus/shrtnr/server/storage/url"
	"sync"
)

var (
	urls trie.Trie
	lock sync.RWMutex
	ctx  context.Context

	repository Repository
)

func Start(parentCtx context.Context, maxUrl uint, maxConcurrency uint) {
	ctx = parentCtx
	urls = trie.NewTrie(maxUrl, maxConcurrency)
	repository = storage.GetURLRepository()

	created.Subscribe(onUrlCreated)
	deleted.Subscribe(onUrlDeleted)
}

func onUrlCreated(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	return errors.New("search: implement onUrlCreated")
}

func onUrlDeleted(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	return errors.New("search: implement onUrlDeleted")
}
