package find

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type getURLResponse struct {
	record entities.URL
	id     entities.ID
	err    error
}

func (s getURLResponse) Success() bool {
	return s.err == nil
}

func buildGetURLResponse(id entities.ID) getURLResponse {
	return getURLResponse{id: id}
}

func getURL(response getURLResponse) getURLResponse {
	response.record, response.err = repository.GetByID(response.id)

	return response
}
