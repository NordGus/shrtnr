package search

import (
	"errors"
	"strings"

	"github.com/NordGus/shrtnr/domain/url/entities"
)

func onUrlDeletedSubscriber(record entities.URL) error {
	lock.Lock()
	defer lock.Unlock()

	clearTargetEntry := strings.TrimPrefix("https://", record.Target.String())
	clearTargetEntry = strings.TrimPrefix("http://", clearTargetEntry)

	err := errors.Join(
		clearTargetCache.RemoveEntry(clearTargetEntry),
		fullTargetCache.RemoveEntry(record.Target.String()),
		shortCache.RemoveEntry(record.UUID.String()),
	)

	return err
}
