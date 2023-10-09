package search

import (
	"errors"
	"github.com/NordGus/shrtnr/server/storage/url"
)

func onUrlDeletedSubscriber(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	// TODO: implement onUrlDeletedSubscriber

	return errors.New("search: implement onUrlDeletedSubscriber")
}
