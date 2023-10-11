package inmemory

import (
	"errors"
	"strings"
	"time"
)

var (
	RecordNotFoundErr = errors.New("inmemory: record not found")
)

type InitFunc[T Record] func(id uint, uuid string, fullURL string, createdAt time.Time) T

type DeletedAtFunc[T Record] func(record T, at time.Time) T

type Storage[T Record] struct {
	initFunc         InitFunc[T]
	setDeletedAtFunc DeletedAtFunc[T]
	store            string
}

func NewInMemoryStorage[T Record](table string, initFunc InitFunc[T], setDeletedAtFunc DeletedAtFunc[T]) *Storage[T] {
	_, ok := store[table]
	if !ok {
		store[table] = Table{
			records:   make([]Record, 0, 10),
			currentID: 1,
		}
	}

	return &Storage[T]{
		initFunc:         initFunc,
		setDeletedAtFunc: setDeletedAtFunc,
		store:            table,
	}
}

func (s *Storage[T]) GetByShort(short string) (T, error) {
	var record T

	lock.RLock()
	defer lock.RUnlock()

	for _, r := range store[s.store].records {
		if r.UUID() == short {
			return r.(T), nil
		}
	}

	return record, RecordNotFoundErr
}

func (s *Storage[T]) GetByFull(full string) (T, error) {
	lock.RLock()
	defer lock.RUnlock()

	var record T

	for _, r := range store[s.store].records {
		if r.FullURL() == full {
			return r.(T), nil
		}
	}

	return record, RecordNotFoundErr
}

func (s *Storage[T]) CreateURL(short string, full string) (T, error) {
	lock.Lock()
	defer lock.Unlock()

	var (
		record = s.initFunc(store[s.store].currentID, short, full, time.Now())
		table  = store[s.store]
	)

	table.records = append(table.records, record)
	table.currentID++
	store[s.store] = table

	return record, nil
}

func (s *Storage[T]) DeleteURL(id uint) (T, error) {
	lock.Lock()
	defer lock.Unlock()

	var (
		record     T
		newRecords = make([]Record, 0, len(store[s.store].records))
		table      = store[s.store]
	)

	for _, u := range table.records {
		if u.ID() != id {
			newRecords = append(newRecords, u)
		} else {
			record = u.(T)
		}
	}

	s.setDeletedAtFunc(record, time.Now())

	if record.ID() == 0 {
		return record, RecordNotFoundErr
	}

	table.records = newRecords
	store[s.store] = table

	return record, nil
}

func (s *Storage[T]) GetLikeLongs(linkLongs ...string) ([]T, error) {
	lock.Lock()
	defer lock.Unlock()

	var records = make([]T, 0, len(store[s.store].records))

	// super inefficient search
	for _, record := range store[s.store].records {
		for _, long := range linkLongs {
			if strings.Contains(record.FullURL(), long) {
				records = append(records, record.(T))
				break
			}
		}
	}

	return records, nil
}
