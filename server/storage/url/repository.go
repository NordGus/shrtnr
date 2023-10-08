package url

import (
	"github.com/NordGus/shrtnr/server/storage"
	"github.com/NordGus/shrtnr/server/storage/url/inmemory"
)

func NewRepository(env string) storage.URLRepository {
	switch env {
	case "production":
		return inmemory.NewInMemoryStorage() // TODO: change for a Database storage when implemented
	case "test":
		return inmemory.NewInMemoryStorage()
	default:
		return inmemory.NewInMemoryStorage()
	}
}
