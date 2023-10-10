package storage

import (
	"github.com/NordGus/shrtnr/domain/storage/url"
)

var (
	urlRepository URLRepository
)

func Start(env string) error {
	if err := initialize(env); err != nil {
		return err
	}

	urlRepository = url.NewRepository(env)

	return nil
}

func GetURLRepository() URLRepository {
	return urlRepository
}
