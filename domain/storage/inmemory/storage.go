package inmemory

import (
	"errors"
	"strings"
	"sync"
	"time"
)

var (
	RecordNotFoundErr = errors.New("inmemory: record not found")

	store map[string][]Record
	lock  sync.RWMutex
)

type InitFunc[T Record] func(id uint, uuid string, fullURL string, createdAt time.Time) T

type DeletedAtFunc[T Record] func(record T, at time.Time) T

type Storage[T Record] struct {
	initFunc         InitFunc[T]
	setDeletedAtFunc DeletedAtFunc[T]
	store            string
	currentID        uint
}

func Start() error {
	store = make(map[string][]Record, 10)

	return nil
}

func NewInMemoryStorage[T Record](table string, initFunc InitFunc[T], setDeletedAtFunc DeletedAtFunc[T]) *Storage[T] {
	_, ok := store[table]
	if !ok {
		store[table] = make([]Record, 0, 10)
	}

	return &Storage[T]{
		initFunc:         initFunc,
		setDeletedAtFunc: setDeletedAtFunc,
		store:            table,
		currentID:        1,
	}
}

func (s *Storage[T]) GetByShort(short string) (T, error) {
	var record T

	lock.RLock()
	defer lock.RUnlock()

	for _, r := range store[s.store] {
		if r.UUID() == short {
			return r.(T), nil
		}
	}

	return record, RecordNotFoundErr
}

func (s *Storage[T]) GetByFull(full string) (T, error) {
	var record T

	lock.RLock()
	defer lock.RUnlock()

	for _, r := range store[s.store] {
		if r.FullURL() == full {
			return r.(T), nil
		}
	}

	return record, RecordNotFoundErr
}

func (s *Storage[T]) CreateURL(short string, full string) (T, error) {
	lock.Lock()
	defer lock.Unlock()

	record := s.initFunc(s.currentID, short, full, time.Now())

	store[s.store] = append(store[s.store], record)
	s.currentID++

	return record, nil
}

func (s *Storage[T]) DeleteURL(short string) (T, error) {
	lock.Lock()
	defer lock.Unlock()

	var (
		record     T
		newRecords = make([]Record, 0, len(store[s.store]))
	)

	for _, u := range store[s.store] {
		if u.UUID() != short {
			newRecords = append(newRecords, u)
		} else {
			record = u.(T)
		}
	}

	s.setDeletedAtFunc(record, time.Now())

	if record.ID() == 0 {
		return record, RecordNotFoundErr
	}

	return record, nil
}

func (s *Storage[T]) GetLikeLongs(linkLongs ...string) ([]T, error) {
	lock.Lock()
	defer lock.Unlock()

	var records = make([]T, 0, len(store[s.store]))

	// super inefficient search
	for _, record := range store[s.store] {
		for _, long := range linkLongs {
			if strings.Contains(record.FullURL(), long) {
				records = append(records, record.(T))
				break
			}
		}
	}

	return records, nil
}
