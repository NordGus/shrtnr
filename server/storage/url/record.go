package url

import "time"

type URL struct {
	Id        uint
	UUID      string
	FullURL   string
	CreatedAt time.Time
	DeletedAt time.Time
}

func (u URL) ID() uint {
	return u.Id
}

func (u URL) Short() string {
	return u.UUID
}

func (u URL) Full() string {
	return u.FullURL
}
