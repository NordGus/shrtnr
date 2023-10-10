package storage

import (
	"github.com/NordGus/shrtnr/domain/shared/response"
	"github.com/NordGus/shrtnr/domain/storage/inmemory"
)

type signal struct {
	env string
	err error
}

func (s signal) Error() error {
	return s.err
}

func initialize(env string) error {
	sig := signal{env: env}

	response.AndThen(sig, startStorage)

	return sig.err
}

func startStorage(sig signal) signal {
	switch sig.env {
	case "production":
		sig.err = inmemory.Start()
	case "test":
		sig.err = inmemory.Start()
	default:
		sig.err = inmemory.Start()
	}

	return sig
}
