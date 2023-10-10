package url

import (
	"time"
)

type URL struct {
	Id           uint
	Uuid         string
	FullUrl      string
	CreationTime time.Time
	DeletionTime time.Time
}

func (u URL) ID() uint {
	return u.Id
}

func (u URL) UUID() string {
	return u.Uuid
}

func (u URL) FullURL() string {
	return u.FullUrl
}

func (u URL) CreatedAt() time.Time {
	return u.CreationTime
}

func (u URL) DeletedAt() time.Time {
	return u.DeletionTime
}

func setURLDeletedAt(record URL, at time.Time) URL {
	record.DeletionTime = at

	return record
}
