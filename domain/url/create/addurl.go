package create

import (
	"errors"

	"github.com/NordGus/shrtnr/domain/shared/queue"
	"github.com/NordGus/shrtnr/domain/url/entities"
	"github.com/NordGus/shrtnr/domain/url/messagebus/created"
	"github.com/NordGus/shrtnr/domain/url/messagebus/deleted"
)

var (
	NonUniqueUUIDErr   = errors.New("create: duplicated UUID")
	NonUniqueTargetErr = errors.New("create: duplicated Target")
)

type addURLResponse struct {
	new       entities.URL
	oldRecord entities.URL
	record    entities.URL
	err       error
}

func (s addURLResponse) Success() bool {
	return s.err == nil
}

func newAddURLResponse(entity entities.URL) addURLResponse {
	return addURLResponse{new: entity, oldRecord: entities.URL{}, record: entities.URL{}, err: nil}
}

func validateUUIDUniqueness(response addURLResponse) addURLResponse {
	_, err := repository.GetByUUID(response.new.UUID)
	if err == nil {
		response.err = errors.Join(response.err, NonUniqueUUIDErr)
	}

	return response
}

func validateTargetUniqueness(response addURLResponse) addURLResponse {
	_, err := repository.GetByTarget(response.new.Target)
	if err == nil {
		response.err = errors.Join(response.err, NonUniqueTargetErr)
	}

	return response
}

func canBeAdded(response addURLResponse) addURLResponse {
	lock.RLock()
	defer lock.RUnlock()

	if cache.IsFull() {
		response.oldRecord, _ = cache.Peek()
		response.err = errors.Join(response.err, queue.IsFullErr)
	}

	return response
}

func deleteOldestUrl(response addURLResponse) addURLResponse {
	if !errors.Is(response.err, queue.IsFullErr) {
		return response
	}

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

	err = created.Raise(record)
	if err != nil {
		response.err = errors.Join(response.err, err)

		return response
	}

	response.record = record

	return response
}

func addUrlToQueue(response addURLResponse) addURLResponse {
	lock.Lock()
	defer lock.Unlock()

	if cache.IsFull() {
		_, err := cache.Pop() // ignores the popped record because the addURLResponse already contains it from the deletion part
		response.err = errors.Join(response.err, err)
	}

	response.err = errors.Join(response.err, cache.Push(response.record))

	return response
}
