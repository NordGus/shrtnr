package create

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/messagebus/url/created"
	"github.com/NordGus/shrtnr/domain/messagebus/url/deleted"
	"github.com/NordGus/shrtnr/domain/shared/queue"
	"github.com/NordGus/shrtnr/domain/storage/url"
)

type signal struct {
	new       URL
	oldRecord url.URL
	record    url.URL
	err       error
}

func (s signal) Error() error {
	return s.err
}

func buildUrl(short string, full string) signal {
	return signal{new: URL{short: shortURL(short), full: fullURL(full)}}
}

func validateUrl(sig signal) signal {
	sig.err = sig.new.Validate()

	return sig
}

func canBeAdded(sig signal) signal {
	if cache.IsFull() {
		sig.oldRecord, _ = cache.Peek()
		sig.err = queue.IsFullErr
	}

	return sig
}

func deleteOldestUrl(sig signal) signal {
	record, err := repository.DeleteURL(sig.oldRecord.ID())
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	err = deleted.Raise(record)
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	return signal{new: sig.new, oldRecord: sig.oldRecord}
}

func persistNewURl(sig signal) signal {
	sig.record, sig.err = repository.CreateURL(string(sig.new.short), string(sig.new.full))
	if sig.err != nil {
		return sig
	}

	sig.err = created.Raise(sig.record)
	if sig.err != nil {
		return sig
	}

	return sig
}

func addUrlToQueue(sig signal) signal {
	if cache.IsFull() {
		_, sig.err = cache.Pop() // ignores the popped record because the signal already contains it from the deletion parte
	}

	sig.err = cache.Push(sig.record)

	return sig
}
