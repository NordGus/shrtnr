package search

import (
	"strings"

	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

func onUrlDeletedSubscriber(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	entry := strings.TrimPrefix("https://", record.FullUrl)
	entry = strings.TrimPrefix("http://", entry)

	return cache.RemoveEntry(entry)
}
