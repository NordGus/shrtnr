package create

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

func onUrlDeletedSubscriber(record entities.URL) error {
	lock.Lock()
	defer lock.Unlock()

	// Ignoring error because there's a possibility that it was popped from the queue before this execution
	_ = cache.FindAndRemoveBy(record, func(i entities.URL, j entities.URL) bool { return i.ID == j.ID })

	return nil
}
