package storage

import "github.com/NordGus/rom-stack/server/storage/url"

var (
	urlRepository URLRepository
)

func Start(env string) {
	urlRepository = url.NewRepository(env)
}

func GetURLRepository() URLRepository {
	return urlRepository
}
