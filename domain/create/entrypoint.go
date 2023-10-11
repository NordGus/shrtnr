package create

import (
	"context"
	"github.com/NordGus/shrtnr/domain/shared/queue"
	"github.com/NordGus/shrtnr/domain/storage"
	"github.com/NordGus/shrtnr/domain/storage/url"
	"sync"
)

var (
	ctx        context.Context
	cache      queue.Queue[url.URL]
	repository Repository

	lock sync.Mutex
)

func Start(otherCtx context.Context, maxUrl uint) {
	ctx = otherCtx
	cache = queue.NewQueue[url.URL](maxUrl)
	repository = storage.GetURLRepository()
}
