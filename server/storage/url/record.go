package url

import "time"

type URL struct {
	Id        uint
	UUID      string
	FullURL   string
	CreatedAt time.Time
	DeletedAt time.Time
}
