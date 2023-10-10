package inmemory

import (
	"errors"
	"strings"
	"sync"
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
	records          []T
	currentID        uint
	lock             sync.RWMutex
}

func NewInMemoryStorage[T Record](initFunc InitFunc[T], setDeletedAtFunc DeletedAtFunc[T]) *Storage[T] {
	return &Storage[T]{
		initFunc:         initFunc,
		setDeletedAtFunc: setDeletedAtFunc,
		records:          make([]T, 0, 100),
		currentID:        1,
	}
}

func (s *Storage[T]) GetByShort(short string) (T, error) {
	var record T

	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, r := range s.records {
		if r.UUID() == short {
			return r, nil
		}
	}

	return record, RecordNotFoundErr
}

func (s *Storage[T]) GetByFull(full string) (T, error) {
	var record T

	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, r := range s.records {
		if r.FullURL() == full {
			return record, nil
		}
	}

	return record, RecordNotFoundErr
}

func (s *Storage[T]) CreateURL(short string, full string) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	record := s.initFunc(s.currentID, short, full, time.Now())

	s.records = append(s.records, record)
	s.currentID++

	return record, nil
}

func (s *Storage[T]) DeleteURL(short string) (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var (
		record     T
		newRecords = make([]T, 0, len(s.records))
	)

	for _, u := range s.records {
		if u.UUID() != short {
			newRecords = append(newRecords, u)
		} else {
			record = u
		}
	}

	s.setDeletedAtFunc(record, time.Now())

	if record.ID() == 0 {
		return record, RecordNotFoundErr
	}

	return record, nil
}

func (s *Storage[T]) GetLikeLongs(linkLongs ...string) ([]T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var records = make([]T, 0, len(s.records))

	// super inefficient search
	for _, record := range s.records {
		for _, long := range linkLongs {
			if strings.Contains(record.FullURL(), long) {
				records = append(records, record)
				break
			}
		}
	}

	return records, nil
}
