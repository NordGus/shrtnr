package create

import (
	"context"
	"sync"

	"github.com/NordGus/shrtnr/domain/shared/queue"
	"github.com/NordGus/shrtnr/domain/url/storage"
	"github.com/NordGus/shrtnr/domain/url/storage/url"
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
