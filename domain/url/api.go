package url

import (
	"github.com/NordGus/shrtnr/domain/url/create"
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
}

func PaginateURLs[T Response[T]](page uint, resp []T) ([]T, int, error) {
	records, err := find.PaginateURLs(page, uint(len(resp)))
	if err != nil {
		return resp, len(records), err
	}

	for i := 0; i < len(records); i++ {
		record := records[i]

		resp[i] = resp[i].SetID(record.ID.String()).
			SetUUID(record.UUID.String()).
			SetTarget(record.Target.String()).
			SetCreatedAt(record.CreatedAt.Time())
	}

	return resp, len(records), nil
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
		SetCreatedAt(record.CreatedAt.Time())

	return resp, nil
}

func CreateURL[T Response[T]](id string, uuid string, target string, resp T) (T, error) {
	record, err := entities.NewURL(id, uuid, target, time.Now())
	if err != nil {
		return resp, err
	}

	record, err = create.AddURL(record)
	if err != nil {
		return resp, err
	}

	resp = resp.SetID(record.ID.String()).
		SetUUID(record.UUID.String()).
		SetTarget(record.Target.String()).
		SetCreatedAt(record.CreatedAt.Time())

	return resp, nil
}
