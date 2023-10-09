package ingest

import (
	"errors"
	"github.com/NordGus/shrtnr/server/messagebus/url/created"
	"github.com/NordGus/shrtnr/server/messagebus/url/deleted"
	"github.com/NordGus/shrtnr/server/shared/queue"
	"github.com/NordGus/shrtnr/server/shared/response"
)

type signal struct {
	new URL
	old URL
	err error
}

func (s signal) Error() error {
	return s.err
}

func AddURL(short string, full string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		lock.Lock()
		defer lock.Unlock()

		resp := response.AndThen(buildUrl(short, full), validateUrl)
		resp = response.AndThen(resp, canBeAdded)
		resp = response.OnFailure(resp, deleteOldestUrl)
		resp = response.AndThen(resp, persistNewURl)
		resp = response.AndThen(resp, addUrlToQueue)

		return resp.err
	}
}

func buildUrl(short string, full string) signal {
	return signal{
		new: URL{short: shortURL(short), full: fullURL(full)},
		old: URL{},
		err: nil,
	}
}

func validateUrl(sig signal) signal {
	sig.err = sig.new.Validate()

	return sig
}

func canBeAdded(sig signal) signal {
	if urls.IsFull() {
		sig.old, _ = urls.Peek()
		sig.err = queue.IsFullErr
	}

	return sig
}

func deleteOldestUrl(sig signal) signal {
	record, err := repository.DeleteURL(string(sig.old.short))
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	err = deleted.Raise(record.Short())
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	return signal{new: sig.new}
}

func persistNewURl(sig signal) signal {
	record, err := repository.CreateURL(string(sig.new.short), string(sig.new.full))
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	err = created.Raise(record.Short())
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	return sig
}

func addUrlToQueue(sig signal) signal {
	if urls.IsFull() {
		sig.old, _ = urls.Pop()
	}

	_ = urls.Push(sig.new)

	return signal{new: sig.new, old: sig.old}
}
