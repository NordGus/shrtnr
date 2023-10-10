package search

import (
	"github.com/NordGus/shrtnr/server/storage/url"
	"strings"
)

func onUrlCreatedSubscriber(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	entry := strings.TrimPrefix("https://", record.FullUrl)
	entry = strings.TrimPrefix("http://", entry)

	cache.AddEntry(entry)

	return nil
}
