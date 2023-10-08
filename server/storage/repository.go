package storage

import "github.com/NordGus/rom-stack/server/ingest"

type URLRepository interface {
	ingest.Repository
}
