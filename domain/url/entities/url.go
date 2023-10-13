package entities

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/google/uuid"
	"time"
)

// URL is the domain representation of an url entity in the application
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

// ID represents the URL entity's storage uuid.UUID
type ID uuid.UUID

// newID validates the given id and translates it to the domain specific ID
func newID(response newURLResponse) newURLResponse {
	validUUID, err := uuid.Parse(response.id)
	if err != nil {
		response.err = errors.Join(response.err, err)
	}

	if response.err == nil {
		response.record.ID = ID(validUUID)
	}

	return response
}

type UUID string

// NewUUID validates the given uuid and translates it to the domain specific UUID
func NewUUID(uuid string) (UUID, error) {
	var out UUID

	if len(uuid) != 8 {
		return out, errors.New("url: uuid too long")
	}

	out = UUID(uuid)

	return out, nil
}

// newUUID validates the given uuid and translates it to the domain specific UUID
func newUUID(response newURLResponse) newURLResponse {
	u, err := NewUUID(response.uuid)
	if err != nil {
		response.err = errors.Join(response.err, err)
	}

	if response.err == nil {
		response.record.UUID = u
	}

	return response
}

type Target string

// newTarget validates the given target and translates it to the domain specific Target
func newTarget(response newURLResponse) newURLResponse {
	// TODO: check if the link is reachable

	if response.err == nil {
		response.record.Target = Target(response.target)
	}

	return response
}

func (t Target) String() string {
	return string(t)
}

type CreatedAt time.Time

// newCreatedAt validates the given createdAt and translates it to the domain specific CreatedAt
func newCreatedAt(response newURLResponse) newURLResponse {
	if response.createdAt.After(time.Now()) {
		response.err = errors.Join(response.err, errors.New("url: can't be created in the future"))
	}

	if response.err == nil {
		response.record.CreatedAt = CreatedAt(response.createdAt)
	}

	return response
}

type DeletedAt time.Time

// newDeletedAt validates the given deletedAt and translates it to the domain specific DeletedAt
func newDeletedAt(response newURLResponse) newURLResponse {
	var comp time.Time

	if response.deletedAt.Before(response.createdAt) && !response.deletedAt.Equal(comp) {
		response.err = errors.Join(response.err, errors.New("url: can't be deleted before being created"))
	}

	if response.err == nil {
		response.record.DeletedAt = DeletedAt(response.deletedAt)
	}

	return response
}
