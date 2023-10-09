package url

import (
	"errors"
	"sync"
	"time"
)

var (
	RecordNotFoundErr = errors.New("url: in memory: record not found")
)

type InMemoryStorage struct {
	records   []URL
	currentID uint
	lock      sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		records:   make([]URL, 0, 100),
		currentID: 1,
	}
}

func (s *InMemoryStorage) GetByShort(short string) (URL, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, record := range s.records {
		if record.UUID == short {
			return record, nil
		}
	}

	return URL{}, RecordNotFoundErr
}

func (s *InMemoryStorage) GetByFull(full string) (URL, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, record := range s.records {
		if record.FullURL == full {
			return record, nil
		}
	}

	return URL{}, RecordNotFoundErr
}

func (s *InMemoryStorage) CreateURL(short string, full string) (URL, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	record := URL{
		Id:        s.currentID,
		UUID:      short,
		FullURL:   full,
		CreatedAt: time.Time{},
		DeletedAt: time.Time{},
	}

	s.records = append(s.records, record)
	s.currentID++

	return record, nil
}

func (s *InMemoryStorage) DeleteURL(short string) (URL, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var (
		record     URL
		newRecords = make([]URL, 0, len(s.records))
	)

	for _, u := range s.records {
		if u.UUID != short {
			newRecords = append(newRecords, u)
		} else {
			record = u
		}
	}

	if record.Id == 0 {
		return record, RecordNotFoundErr
	}

	return record, nil
}
