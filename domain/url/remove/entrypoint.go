package remove

import (
	"context"
	"sync"

	"github.com/NordGus/shrtnr/domain/url/storage"
)

var (
	ctx        context.Context
	repository Repository

	lock sync.Mutex
)

func Start(otherCtx context.Context) {
	ctx = otherCtx
	repository = storage.GetRepository()
}
