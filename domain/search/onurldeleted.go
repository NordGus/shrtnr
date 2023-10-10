package search

import (
	"github.com/NordGus/shrtnr/domain/storage/url"
	"strings"
)

func onUrlDeletedSubscriber(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	entry := strings.TrimPrefix("https://", record.FullUrl)
	entry = strings.TrimPrefix("http://", entry)

	return cache.RemoveEntry(entry)
}
