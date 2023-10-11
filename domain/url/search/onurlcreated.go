package search

import (
	"strings"

	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

func onUrlCreatedSubscriber(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	entry := strings.TrimPrefix("https://", record.FullUrl)
	entry = strings.TrimPrefix("http://", entry)

	cache.AddEntry(entry)

	return nil
}
