package create

import (
	"context"
	"log"
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

	fillCache(maxUrl)
}

func fillCache(recordsLimit uint) {
	rcrds, err := repository.GetAllRecords(recordsLimit)
	if err != nil {
		log.Println(err) // ignores errors because there can be no records and still return an error
	}

	for _, rcrd := range rcrds {
		// ignoring error because in the worse case scenario the number of records is equal to the max size of the queue cache
		_ = cache.Push(rcrd)
	}
}
