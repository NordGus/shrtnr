package search

import (
	"errors"
	"log"
	"strings"

	"github.com/NordGus/shrtnr/domain/url/entities"
)

func onUrlDeletedSubscriber(record entities.URL) error {
	lock.Lock()
	defer lock.Unlock()

	clearTargetEntry := strings.TrimPrefix(record.Target.String(), "https://")
	clearTargetEntry = strings.TrimPrefix(clearTargetEntry, "http://")
	clearTargetEntry = strings.TrimPrefix(clearTargetEntry, "www.")

	err1 := clearTargetCache.RemoveEntry(clearTargetEntry)
	if err1 != nil {
		log.Println("search: failed to remove entry from clearTargetCache", err1)
	}

	err2 := fullTargetCache.RemoveEntry(record.Target.String())
	if err2 != nil {
		log.Println("search: failed to remove entry from fullTargetCache", err2)
	}

	err3 := shortCache.RemoveEntry(record.UUID.String())
	if err3 != nil {
		log.Println("search: failed to remove entry from shortCache", err3)
	}

	return errors.Join(err1, err2, err3)
}
