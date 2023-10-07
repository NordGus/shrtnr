package ingestion

import (
	"errors"
	"log"
)

var (
	ErrIngestionAbnormalFailure = errors.New("ingestion: something went horribly wrong while ingesting")
)

func AddURL(short string, full string) error {
	lock.Lock()
	defer lock.Unlock()

	record := Url{short: short, full: full}

	if urls.Push(record) != nil {
		_, _ = urls.Pop()
		if urls.Push(record) != nil {
			log.Fatalln(ErrIngestionAbnormalFailure, ": url", short, full)
		}
	}

	// TODO implement creation flow

	return nil
}
