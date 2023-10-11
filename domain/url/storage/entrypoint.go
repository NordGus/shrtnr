package storage

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

var (
	urlRepository URLRepository
)

func Start(env string) error {
	urlRepository = url.NewRepository(env)

	return nil
}

func GetURLRepository() URLRepository {
	return urlRepository
}