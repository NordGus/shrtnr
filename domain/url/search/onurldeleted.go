package search

import (
	"strings"

	"github.com/NordGus/shrtnr/domain/url/entities"
)

func onUrlDeletedSubscriber(record entities.URL) error {
	lock.Lock()
	defer lock.Unlock()

	entry := strings.TrimPrefix("https://", record.Target.String())
	entry = strings.TrimPrefix("http://", entry)

	return cache.RemoveEntry(entry)
}
