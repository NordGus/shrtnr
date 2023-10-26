package url

import (
	"github.com/NordGus/shrtnr/domain/url/create"
	"github.com/NordGus/shrtnr/domain/url/entities"
	"github.com/NordGus/shrtnr/domain/url/find"
	"github.com/NordGus/shrtnr/domain/url/remove"
	"github.com/NordGus/shrtnr/domain/url/search"
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

func CreateURL[T Response[T]](id string, uuid string, target string, resp T, oldResp T) (T, T, error) {
	var (
		err       error
		record    entities.URL
		oldRecord entities.URL
	)

	record, err = entities.NewURL(id, uuid, target, time.Now())
	if err != nil {
		return resp, oldResp, err
	}

	record, oldRecord, err = create.AddURL(record)
	if err != nil {
		return resp, oldResp, err
	}

	resp = resp.SetID(record.ID.String()).
		SetUUID(record.UUID.String()).
		SetTarget(record.Target.String()).
		SetCreatedAt(record.CreatedAt.Time())

	oldResp = oldResp.SetID(oldRecord.ID.String()).
		SetUUID(oldRecord.UUID.String()).
		SetTarget(oldRecord.Target.String()).
		SetCreatedAt(oldRecord.CreatedAt.Time())

	return resp, oldResp, err
}

func RemoveURL[T Response[T]](id string, resp T) (T, error) {
	var (
		err      error
		recordID entities.ID
		record   entities.URL
	)

	recordID, err = entities.NewID(id)
	if err != nil {
		return resp, err
	}

	record, err = remove.DeleteURL(recordID)
	if err != nil {
		return resp, err
	}

	resp = resp.SetID(record.ID.String()).
		SetUUID(record.UUID.String()).
		SetTarget(record.Target.String()).
		SetCreatedAt(record.CreatedAt.Time())

	return resp, err
}

func SearchURLsBy[T Response[T]](term string, resp []T) ([]T, error) {
	records, err := search.ByTerm(term)
	if err != nil {
		return resp, err
	}

	resp = make([]T, len(records))

	for i, record := range records {
		resp[i] = resp[i].
			SetID(record.ID.String()).
			SetUUID(record.UUID.String()).
			SetTarget(record.Target.String()).
			SetCreatedAt(record.CreatedAt.Time())
	}

	return resp, nil
}
