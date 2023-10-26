package remove

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/url/messagebus/deleted"
	"log"

	"github.com/NordGus/shrtnr/domain/url/entities"
)

var (
	RecordNotFoundErr          = errors.New("remove: record not found")
	FailedRecordDeletionErr    = errors.New("remove: failed to remove record")
	FailedDeletePropagationErr = errors.New("remove: failed record deletion propagation, panicking")
)

type removeURLResponse struct {
	id     entities.ID
	record entities.URL
	err    error
}

func (s removeURLResponse) Success() bool {
	return s.err == nil
}

func newRemoveURLResponse(id entities.ID) removeURLResponse {
	return removeURLResponse{id: id, record: entities.URL{}, err: nil}
}

func findRecord(response removeURLResponse) removeURLResponse {
	record, err := repository.GetByID(response.id)
	if err != nil {
		response.err = errors.Join(response.err, RecordNotFoundErr, err)
	}

	if err == nil {
		response.record = record
	}

	return response
}

func deleteRecord(response removeURLResponse) removeURLResponse {
	_, err := repository.DeleteURL(response.id)
	if err != nil {
		response.err = errors.Join(response.err, FailedRecordDeletionErr, err)
	}

	return response
}

func propagateDeletion(response removeURLResponse) removeURLResponse {
	err := deleted.Raise(response.record)
	if err != nil {
		// panics because the system has been corrupted
		log.Fatalln(FailedDeletePropagationErr, response.err, err)
	}

	return response
}
