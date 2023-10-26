package search

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/url/entities"
	"log"
)

func onUrlDeletedSubscriber(record entities.URL) error {
	lock.Lock()
	defer lock.Unlock()

	err1 := clearTargetCache.RemoveEntry(clearTargetEntry(record.Target.String()))
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
