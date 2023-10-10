package search

import (
	"github.com/NordGus/shrtnr/server/storage/url"
	"strings"
)

func onUrlDeletedSubscriber(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	entry := strings.TrimPrefix("https://", record.FullURL)
	entry = strings.TrimPrefix("http://", entry)

	return cache.RemoveEntry(entry)
}
