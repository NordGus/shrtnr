package ingest

type Record interface {
	ID() string
	Short() string
	Full() string
}

type Repository interface {
	GetByShort(short string) (Record, error)
	GetByFull(full string) (Record, error)
	CreateURL(short string, full string) (Record, error)
	DeleteURL(short string) (Record, error)
}
