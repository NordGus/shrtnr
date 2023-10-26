package search

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

func onUrlCreatedSubscriber(record entities.URL) error {
	lock.Lock()
	defer lock.Unlock()

	clearTargetCache.AddEntry(clearTargetEntry(record.Target.String()))
	fullTargetCache.AddEntry(record.Target.String())
	shortCache.AddEntry(record.UUID.String())

	return nil
}
