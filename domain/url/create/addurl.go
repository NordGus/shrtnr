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

func canBeAdded(response addURLResponse) addURLResponse {
	if cache.IsFull() {
		response.oldRecord, _ = cache.Peek()
		response.err = queue.IsFullErr
	}

	return response
}

func deleteOldestUrl(response addURLResponse) addURLResponse {
	record, err := repository.DeleteURL(response.oldRecord.ID)
	if err != nil {
		response.err = errors.Join(response.err, err)

		return response
	}

	err = deleted.Raise(record)
	if err != nil {
		response.err = errors.Join(response.err, err)

		return response
	}

	return addURLResponse{new: response.new, oldRecord: record}
}

func persistNewURl(response addURLResponse) addURLResponse {
	record, err := repository.CreateURL(response.new)
	if err != nil {
		response.err = errors.Join(response.err, err)

		return response
	}

	response.err = created.Raise(record)
	if response.err != nil {
		response.err = errors.Join(response.err, err)

		return response
	}

	response.record = record

	return response
}

func addUrlToQueue(response addURLResponse) addURLResponse {
	if cache.IsFull() {
		_, response.err = cache.Pop() // ignores the popped record because the addURLResponse already contains it from the deletion parte
	}

	response.err = cache.Push(response.record)

	return response
}
