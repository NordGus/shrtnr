package find

import (
	"context"

	"github.com/NordGus/shrtnr/domain/url/storage"
)

var (
	ctx        context.Context
	repository Repository
)

func Start(parentCtx context.Context) error {
	ctx = parentCtx
	repository = storage.GetRepository()

	return nil
}
