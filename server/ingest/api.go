package ingest

import (
	"github.com/NordGus/rom-stack/server/shared/response"
)

func AddURL(short string, full string) error {
	lock.Lock()
	defer lock.Unlock()

	resp := response.AndThen(buildUrl(short, full), validateUrl)
	resp = response.AndThen(resp, addUrlToQueue)
	resp = response.OnFailure(resp, deleteOldestUrl)
	resp = response.AndThen(resp, persistNewURl)

	return resp.err
}

func buildUrl(short string, full string) signal {
	return signal{
		new: Url{short: short, full: full},
		old: Url{},
		err: nil,
	}
}

func validateUrl(sig signal) signal {
	// do validation

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
	// delete old from system

	return signal{new: sig.new}
}

func persistNewURl(sig signal) signal {
	// persists new url

	return sig
}
