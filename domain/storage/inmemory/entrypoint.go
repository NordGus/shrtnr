package inmemory

import "sync"

var (
	store map[string]Table
	lock  sync.RWMutex
)

func Start() error {
	store = make(map[string]Table, 10)

	return nil
}
