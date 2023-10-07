package ingest

import (
	"context"
	"github.com/NordGus/rom-stack/server/shared/queue"
	"sync"
)

type Url struct {
	short string
	full  string
}

var (
	limit uint

	ctx  context.Context
	urls queue.Queue[Url]

	lock sync.Mutex
)

func Start(otherCtx context.Context, maxUrl uint) {
	limit = maxUrl
	ctx = otherCtx
	urls = queue.NewQueue[Url](maxUrl)
}
