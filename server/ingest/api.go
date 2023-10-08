package ingest

import (
	"errors"
	"github.com/NordGus/rom-stack/server/messagebus/url/created"
	"github.com/NordGus/rom-stack/server/messagebus/url/deleted"
	"github.com/NordGus/rom-stack/server/shared/response"
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
		resp = response.AndThen(resp, addUrlToQueue)
		resp = response.OnFailure(resp, deleteOldestUrl)
		resp = response.AndThen(resp, persistNewURl)

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

func addUrlToQueue(sig signal) signal {
	err := urls.Push(sig.new)
	if err != nil {
		sig.err = err
		sig.old, _ = urls.Pop()
	}

	return sig
}

func deleteOldestUrl(sig signal) signal {
	record, err := repository.DeleteURL(string(sig.old.short))
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	err = deleted.Raise(record.UUID)
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

	err = created.Raise(record.UUID)
	if err != nil {
		sig.err = errors.Join(sig.err, err)

		return sig
	}

	return sig
}
