package deleted

import (
	"errors"
	"sync"

	"github.com/NordGus/shrtnr/domain/url/entities"
)

// Subscribe adds a new subscriber to the event
func Subscribe(sub Subscriber) {
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
func Raise(record entities.URL) error {
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
			go func(wg *sync.WaitGroup, out chan<- error, sub Subscriber, record entities.URL) {
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
