package url

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
	"github.com/NordGus/shrtnr/domain/url/find"
	"time"
)

// Response interface represents entities.URL outside the domain.
type Response[T any] interface {
	SetID(id string) T
	SetUUID(uuid string) T
	SetTarget(target string) T
	SetCreatedAt(createdAt time.Time) T
	SetDeletedAt(deletedAt time.Time) T
}

func FindURLByUUID[T Response[T]](uuid string, resp T) (T, error) {
	id, err := entities.NewUUID(uuid)
	if err != nil {
		return resp, err
	}

	record, err := find.GetByUUID(id)
	if err != nil {
		return resp, err
	}

	resp = resp.SetID(record.ID.String()).
		SetUUID(record.UUID.String()).
		SetTarget(record.Target.String()).
		SetCreatedAt(record.CreatedAt.Time()).
		SetDeletedAt(record.DeletedAt.Time())

	return resp, nil
}
