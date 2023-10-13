package url

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/google/uuid"
	"time"
)

// URL is the domain representation of an url record in the application
type URL struct {
	ID        ID
	UUID      UUID
	Target    Target
	CreatedAt CreatedAt
	DeletedAt DeletedAt
}

// newURLResponse represents the inputs and outputs of the control flow of the newURL function
type newURLResponse struct {
	id        string
	uuid      string
	target    string
	createdAt time.Time
	deletedAt time.Time
	record    URL
	err       error
}

// Success indicates if the newURLResponse was successful
func (s newURLResponse) Success() bool {
	return s.err == nil
}

// newURL translate external data into the domain specific URL struct or returns an error
func newURL(id string, uuid string, target string, createdAt time.Time, deletedAt time.Time) (URL, error) {
	var sig = newURLResponse{id: id, uuid: uuid, target: target, createdAt: createdAt, deletedAt: deletedAt}

	resp := railway.OrThen(sig, newID)
	resp = railway.OrThen(resp, newUUID)
	resp = railway.OrThen(resp, newTarget)
	resp = railway.OrThen(resp, newCreatedAt)
	resp = railway.OrThen(resp, newDeletedAt)

	return sig.record, sig.err
}

type ID string

// newID validates the given id and translates it to the domain specific ID
func newID(sig newURLResponse) newURLResponse {
	_, err := uuid.Parse(sig.id)
	if err != nil {
		sig.err = errors.Join(sig.err, err)
	}

	if sig.err == nil {
		sig.record.ID = ID(sig.id)
	}

	return sig
}

type UUID string

// newUUID validates the given uuid and translates it to the domain specific UUID
func newUUID(sig newURLResponse) newURLResponse {
	if len(sig.uuid) != 8 {
		sig.err = errors.Join(sig.err, errors.New("url: uuid too long"))
	}

	if sig.err == nil {
		sig.record.UUID = UUID(sig.uuid)
	}

	return sig
}

type Target string

// newTarget validates the given target and translates it to the domain specific Target
func newTarget(sig newURLResponse) newURLResponse {
	// TODO: check if the link is reachable

	if sig.err == nil {
		sig.record.Target = Target(sig.target)
	}

	return sig
}

type CreatedAt time.Time

// newCreatedAt validates the given createdAt and translates it to the domain specific CreatedAt
func newCreatedAt(sig newURLResponse) newURLResponse {
	if sig.createdAt.After(time.Now()) {
		sig.err = errors.Join(sig.err, errors.New("url: can't be created in the future"))
	}

	if sig.err == nil {
		sig.record.CreatedAt = CreatedAt(sig.createdAt)
	}

	return sig
}

type DeletedAt time.Time

// newDeletedAt validates the given deletedAt and translates it to the domain specific DeletedAt
func newDeletedAt(sig newURLResponse) newURLResponse {
	var comp time.Time

	if sig.deletedAt.Before(sig.createdAt) && !sig.deletedAt.Equal(comp) {
		sig.err = errors.Join(sig.err, errors.New("url: can't be deleted before being created"))
	}

	if sig.err == nil {
		sig.record.DeletedAt = DeletedAt(sig.deletedAt)
	}

	return sig
}
