package inmemory

import (
	"errors"
	"github.com/NordGus/shrtnr/server/storage/url"
	"sync"
	"time"
)

var (
	RecordNotFoundErr = errors.New("url: in memory: record not found")
)

type Storage struct {
	records   []url.URL
	currentID uint
	lock      sync.RWMutex
}

func NewInMemoryStorage() *Storage {
	return &Storage{
		records:   make([]url.URL, 0, 100),
		currentID: 1,
	}
}

func (s *Storage) GetByShort(short string) (url.URL, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, record := range s.records {
		if record.UUID == short {
			return record, nil
		}
	}

	return url.URL{}, RecordNotFoundErr
}

func (s *Storage) GetByFull(full string) (url.URL, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, record := range s.records {
		if record.FullURL == full {
			return record, nil
		}
	}

	return url.URL{}, RecordNotFoundErr
}

func (s *Storage) CreateURL(short string, full string) (url.URL, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	record := url.URL{
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

func (s *Storage) DeleteURL(short string) (url.URL, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	record, err := s.GetByShort(short)
	if err != nil {
		return record, err
	}

	newRecords := make([]url.URL, 0, len(s.records))

	for _, u := range s.records {
		if u.Id != record.Id {
			newRecords = append(newRecords, u)
		}
	}

	return record, nil
}
