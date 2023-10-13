package find

import (
	"github.com/NordGus/shrtnr/domain/url"
)

type getURLResponse struct {
	record url.URL
	id     url.ID
	err    error
}

func (s getURLResponse) Success() bool {
	return s.err == nil
}

func buildGetSignal(id url.ID) getURLResponse {
	return getURLResponse{id: id}
}

func getURL(sig getURLResponse) getURLResponse {
	sig.record, sig.err = repository.GetByID(sig.id)

	return sig
}
