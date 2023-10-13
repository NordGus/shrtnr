package create

import (
	"context"
	"sync"

	"github.com/NordGus/shrtnr/domain/shared/queue"
	"github.com/NordGus/shrtnr/domain/url/entities"
	"github.com/NordGus/shrtnr/domain/url/storage"
)

var (
	ctx        context.Context
	cache      queue.Queue[entities.URL]
	repository Repository

	lock sync.Mutex
)

func Start(otherCtx context.Context, maxUrl uint) {
	ctx = otherCtx
	cache = queue.NewQueue[entities.URL](maxUrl)
	repository = storage.GetRepository()
}
