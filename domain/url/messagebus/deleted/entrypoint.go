package deleted

import (
	"context"
	"sync"

	"github.com/NordGus/shrtnr/domain/url/entities"
)

// Subscriber is an alias for the function signature of the message subscribers
type Subscriber func(record entities.URL) error

var (
	ctx         context.Context
	subscribers []Subscriber
	lock        sync.Mutex
)

// Start initializes the deleted message
func Start(parentCtx context.Context) {
	ctx = parentCtx
	subscribers = make([]Subscriber, 0, 10)
}
