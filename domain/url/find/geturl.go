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

func buildGetURLResponse(id url.ID) getURLResponse {
	return getURLResponse{id: id}
}

func getURL(response getURLResponse) getURLResponse {
	response.record, response.err = repository.GetByID(response.id)

	return response
}
