package search

import (
	"strings"

	"github.com/NordGus/shrtnr/domain/url/entities"
)

func onUrlCreatedSubscriber(record entities.URL) error {
	lock.Lock()
	defer lock.Unlock()

	clearTargetEntry := strings.TrimPrefix("https://", record.Target.String())
	clearTargetEntry = strings.TrimPrefix("http://", clearTargetEntry)

	clearTargetCache.AddEntry(clearTargetEntry)
	fullTargetCache.AddEntry(record.Target.String())
	shortCache.AddEntry(record.UUID.String())

	return nil
}
