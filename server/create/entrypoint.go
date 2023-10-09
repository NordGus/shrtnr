package create

import (
	"context"
	"github.com/NordGus/shrtnr/server/shared/queue"
	"github.com/NordGus/shrtnr/server/storage"
	"sync"
)

var (
	ctx        context.Context
	cache      queue.Queue[URL]
	repository Repository

	lock sync.Mutex
)

func Start(otherCtx context.Context, maxUrl uint) {
	ctx = otherCtx
	cache = queue.NewQueue[URL](maxUrl)
	repository = storage.GetURLRepository()
}
