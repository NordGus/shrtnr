package url

import (
	"time"

	"github.com/NordGus/shrtnr/domain/url/storage/inmemory"
)

type Repository interface {
	GetByShort(short string) (URL, error)
	GetByFull(full string) (URL, error)
	CreateURL(short string, full string) (URL, error)
	DeleteURL(id uint) (URL, error)

	GetLikeLongs(likeLongs ...string) ([]URL, error)

	GetByID(id uint) (URL, error)
	GetAllInPage(page uint, perPage uint) ([]URL, error)
}

func NewRepository(env string) Repository {
	switch env {
	case "production":
		return inmemory.NewInMemoryStorage[URL](newURL, setURLDeletedAt) // TODO: change for a Database storage when implemented
	case "test":
		return inmemory.NewInMemoryStorage[URL](newURL, setURLDeletedAt)
	default:
		return inmemory.NewInMemoryStorage[URL](newURL, setURLDeletedAt)
	}
}

func newURL(id uint, uuid string, fullURL string, createdAt time.Time) URL {
	return URL{
		Id:           id,
		Uuid:         uuid,
		FullUrl:      fullURL,
		CreationTime: createdAt,
	}
}
