package find

import (
	"context"
	"github.com/NordGus/shrtnr/domain/storage"
)

var (
	ctx        context.Context
	repository Repository
)

func Start(parentCtx context.Context) error {
	ctx = parentCtx
	repository = storage.GetURLRepository()

	return nil
}
