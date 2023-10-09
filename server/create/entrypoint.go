package create

import (
	"context"
	"github.com/NordGus/shrtnr/server/shared/queue"
	"github.com/NordGus/shrtnr/server/storage"
	"sync"
)

var (
	ctx        context.Context
	urls       queue.Queue[URL]
	repository Repository

	lock sync.Mutex
)

func Start(otherCtx context.Context, maxUrl uint) {
	ctx = otherCtx
	urls = queue.NewQueue[URL](maxUrl)
	repository = storage.GetURLRepository()
}
