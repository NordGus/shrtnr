package created

import (
	"context"
	"sync"

	"github.com/NordGus/shrtnr/domain/url"
)

// Subscriber is an alias for the function signature of the message subscribers
type Subscriber func(record url.URL) error

var (
	ctx         context.Context
	subscribers []Subscriber
	lock        sync.Mutex
)

// Start initializes the created message
func Start(parentCtx context.Context) {
	ctx = parentCtx
	subscribers = make([]Subscriber, 0, 10)
}
