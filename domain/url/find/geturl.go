package find

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

type getSignal struct {
	record url.URL
	id     uint
	err    error
}

func (s getSignal) Error() error {
	return s.err
}

func buildGetSignal(id uint) getSignal {
	return getSignal{id: id}
}

func getURL(sig getSignal) getSignal {
	sig.record, sig.err = repository.GetByID(sig.id)

	return sig
}
