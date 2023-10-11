package deleted

import (
	"context"
	"sync"

	"github.com/NordGus/shrtnr/domain/url/messagebus"
)

var (
	ctx         context.Context
	subscribers []messagebus.Subscriber
	lock        sync.Mutex
)

// Start initializes the deleted message
func Start(parentCtx context.Context) {
	ctx = parentCtx
	subscribers = make([]messagebus.Subscriber, 0, 10)
}
