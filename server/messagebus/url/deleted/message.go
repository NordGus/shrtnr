package deleted

import (
	"context"
	"github.com/NordGus/shrtnr/server/storage/url"
)

type Subscriber func(record url.URL) error

var (
	ctx         context.Context
	subscribers []Subscriber
)

// Start initializes the deleted message
func Start(parentCtx context.Context) {
	ctx = parentCtx
	subscribers = make([]Subscriber, 0, 10)
}

// Subscribe adds a new subscriber to the event
func Subscribe(sub Subscriber) {
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
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		ch := make(chan error)
		innerCtx, cancel := context.WithCancel(ctx)
		defer func(ch chan error, cancel context.CancelFunc) {
			cancel()
			close(ch)
		}(ch, cancel)

		for _, subscriber := range subscribers {
			go func(ctx context.Context, out chan<- error, sub Subscriber, record url.URL) {
				if err := sub(record); err != nil {
					select {
					case <-ctx.Done():
					case out <- err:
					}
				}
			}(innerCtx, ch, subscriber, record)
		}

		for err := range ch {
			return err
		}

		return nil
	}
}
