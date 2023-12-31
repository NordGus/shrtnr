package entities

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/google/uuid"
	neturl "net/url"
	"time"
)

var (
	InvalidUUIDErr         = errors.New("entities: URL uuid invalid")
	TargetBlankErr         = errors.New("entities: URL target can't be blank")
	TargetInvalidFormatErr = errors.New("entities: URL target is not a valid url")
	CreatedInTheFutureErr  = errors.New("entities: URL can't be created in the future")
)

// URL is the domain representation of an url entity in the application
type URL struct {
	ID        ID
	UUID      UUID
	Target    Target
	CreatedAt CreatedAt
}

// newURLResponse represents the inputs and outputs of the control flow of the NewURL function
type newURLResponse struct {
	id        string
	uuid      string
	target    string
	createdAt time.Time
	record    URL
	err       error
}

// Success indicates if the newURLResponse was successful
func (s newURLResponse) Success() bool {
	return s.err == nil
}

// NewURL translate external data into the domain specific URL struct or returns an error
func NewURL(id string, uuid string, target string, createdAt time.Time) (URL, error) {
	var sig = newURLResponse{id: id, uuid: uuid, target: target, createdAt: createdAt, record: URL{}, err: nil}

	resp := railway.Then(sig, newID)
	resp = railway.Then(resp, newUUID)
	resp = railway.Then(resp, newTarget)
	resp = railway.Then(resp, newCreatedAt)

	return resp.record, resp.err
}

// ID represents the URL entity's storage uuid.UUID
type ID uuid.UUID

// NewID validates the given id and translates it to the domain specific ID
func NewID(id string) (ID, error) {
	resp := newURLResponse{id: id}
	resp = railway.AndThen(resp, newID)

	return resp.record.ID, resp.err
}

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

func (i ID) String() string {
	return uuid.UUID(i).String()
}

type UUID string

// NewUUID validates the given uuid and translates it to the domain specific UUID
func NewUUID(uuid string) (UUID, error) {
	resp := newURLResponse{uuid: uuid}
	resp = railway.AndThen(resp, newUUID)

	return resp.record.UUID, resp.err
}

// newUUID validates the given uuid and translates it to the domain specific UUID
func newUUID(response newURLResponse) newURLResponse {
	if len(response.uuid) != 8 {
		response.err = InvalidUUIDErr
	}

	if response.err == nil {
		response.record.UUID = UUID(response.uuid)
	}

	return response
}

func (u UUID) String() string {
	return string(u)
}

type Target string

// newTarget validates the given target and translates it to the domain specific Target
func newTarget(response newURLResponse) newURLResponse {
	if response.target == "" {
		response.err = errors.Join(response.err, TargetBlankErr)
	}

	_, err := neturl.ParseRequestURI(response.target) // validates that the target is at least a valid URI
	if err != nil {
		response.err = errors.Join(response.err, err, TargetInvalidFormatErr)
	}

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
		response.err = errors.Join(response.err, CreatedInTheFutureErr)
	}

	if response.err == nil {
		response.record.CreatedAt = CreatedAt(response.createdAt)
	}

	return response
}

func (c CreatedAt) Time() time.Time {
	return time.Time(c)
}

func (c CreatedAt) Unix() int64 {
	return time.Time(c).Unix()
}

type DeletedAt time.Time
