package storage

import "github.com/NordGus/shrtnr/server/ingest"

type URLRepository interface {
	ingest.Repository
}
