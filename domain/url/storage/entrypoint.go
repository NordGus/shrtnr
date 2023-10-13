package storage

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
	"github.com/NordGus/shrtnr/domain/url/storage/inmemory"
	"github.com/NordGus/shrtnr/domain/url/storage/url"
	"time"
)

var (
	repository Repository
)

func Start(env string) error {
	repository = url.NewRepository(env)

	return nil
}

func GetRepository() Repository {
	return repository
}

func getStorage(env string) Repository {
	switch env {
	case "production":
		return inmemory.NewInMemoryStorage[entities.URL](entities.NewURL, setEntityDeletedAt) // TODO: change for a Database storage when implemented
	case "test":
		return inmemory.NewInMemoryStorage[entities.URL](entities.NewURL, setEntityDeletedAt)
	default:
		return inmemory.NewInMemoryStorage[entities.URL](entities.NewURL, setEntityDeletedAt)
	}
}

func setEntityDeletedAt(entity entities.URL, deletedAt time.Time) entities.URL {

}
