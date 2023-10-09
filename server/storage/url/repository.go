package url

type Repository interface {
	GetByShort(short string) (URL, error)
	GetByFull(full string) (URL, error)
	CreateURL(short string, full string) (URL, error)
	DeleteURL(short string) (URL, error)
	GetLikeLongs(likeLongs ...string) ([]URL, error)
}

func NewRepository(env string) Repository {
	switch env {
	case "production":
		return NewInMemoryStorage() // TODO: change for a Database storage when implemented
	case "test":
		return NewInMemoryStorage()
	default:
		return NewInMemoryStorage()
	}
}
