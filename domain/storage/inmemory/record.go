package inmemory

import "time"

type Record interface {
	ID() uint
	UUID() string
	FullURL() string
	CreatedAt() time.Time
	DeletedAt() time.Time
}
