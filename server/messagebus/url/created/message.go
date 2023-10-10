package created

import (
	"context"
	"errors"
	"github.com/NordGus/shrtnr/server/storage/url"
	"sync"

	urlmsg "github.com/NordGus/shrtnr/server/messagebus/url"
)

var (
	ctx         context.Context
	subscribers []urlmsg.Subscriber
	lock        sync.Mutex
)

// Start initializes the created message
func Start(parentCtx context.Context) {
	ctx = parentCtx
	subscribers = make([]urlmsg.Subscriber, 0, 10)
}

// Subscribe adds a new subscriber to the event
func Subscribe(sub urlmsg.Subscriber) {
	lock.Lock()
	defer lock.Unlock()

	select {
	case <-ctx.Done():
	default:
		subscribers = append(subscribers, sub)
	}
}

// Raise sends the event to every subscriber of the event.
//
// Note: This is completely overengineered but is fun
func Raise(record url.URL) error {
	lock.Lock()
	defer lock.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		var (
			out error
			ch  = make(chan error, len(subscribers))
			wg  = new(sync.WaitGroup)
		)

		wg.Add(len(subscribers))

		for _, subscriber := range subscribers {
			go func(wg *sync.WaitGroup, out chan<- error, sub urlmsg.Subscriber, record url.URL) {
				out <- sub(record)
				wg.Done()
			}(wg, ch, subscriber, record)
		}

		wg.Wait()
		close(ch)

		for err := range ch {
			out = errors.Join(out, err)
		}

		return out
	}
}
