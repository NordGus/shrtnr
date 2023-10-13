package create

import (
	"errors"

	"github.com/NordGus/shrtnr/domain/shared/queue"
	"github.com/NordGus/shrtnr/domain/url"
	"github.com/NordGus/shrtnr/domain/url/messagebus/created"
	"github.com/NordGus/shrtnr/domain/url/messagebus/deleted"
)

type addURLResponse struct {
	new       url.URL
	oldRecord url.URL
	record    url.URL
	err       error
}

func (s addURLResponse) Success() bool {
	return s.err == nil
}

func newAddURLResponse(entity url.URL) addURLResponse {
	return addURLResponse{new: entity}
}

func canBeAdded(sig addURLResponse) addURLResponse {
	if cache.IsFull() {
		sig.oldRecord, _ = cache.Peek()
		sig.err = queue.IsFullErr
	}

	return sig
}

func deleteOldestUrl(sig addURLResponse) addURLResponse {
	record, err := repository.DeleteURL(sig.oldRecord.ID())
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	err = deleted.Raise(record)
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	return addURLResponse{new: sig.new, oldRecord: sig.oldRecord}
}

func persistNewURl(sig addURLResponse) addURLResponse {
	sig.record, sig.err = repository.CreateURL(string(sig.new.short), string(sig.new.full))
	if sig.err != nil {
		return sig
	}

	sig.err = created.Raise(sig.record)
	if sig.err != nil {
		return sig
	}

	return sig
}

func addUrlToQueue(sig addURLResponse) addURLResponse {
	if cache.IsFull() {
		_, sig.err = cache.Pop() // ignores the popped record because the addURLResponse already contains it from the deletion parte
	}

	sig.err = cache.Push(sig.record)

	return sig
}
