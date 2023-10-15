package find

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type getURLByUUIDResponse struct {
	record entities.URL
	uuid   entities.UUID
	err    error
}

func (s getURLByUUIDResponse) Success() bool {
	return s.err == nil
}

func buildGetURLByUUIDResponse(uuid entities.UUID) getURLByUUIDResponse {
	return getURLByUUIDResponse{uuid: uuid}
}

func getURLByUUID(response getURLByUUIDResponse) getURLByUUIDResponse {
	response.record, response.err = repository.GetByUUID(response.uuid)

	return response
}
