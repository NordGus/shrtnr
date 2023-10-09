package search

import (
	"errors"
	"github.com/NordGus/shrtnr/server/storage/url"
)

func onUrlCreatedSubscriber(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	// TODO: implement onUrlCreatedSubscriber

	return errors.New("search: implement onUrlCreatedSubscriber")
}
