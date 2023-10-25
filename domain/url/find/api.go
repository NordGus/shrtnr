package find

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

func PaginateURLs(page uint, perPage uint) ([]entities.URL, error) {
	select {
	case <-ctx.Done():
		return []entities.URL{}, ctx.Err()
	default:
		response := buildPaginateURLsResponse(page, perPage)
		response = railway.AndThen(response, getURLs)

		return response.records, response.err
	}
}

func GetByUUID(uuid entities.UUID) (entities.URL, error) {
	select {
	case <-ctx.Done():
		return entities.URL{}, ctx.Err()
	default:
		response := buildGetURLByUUIDResponse(uuid)
		response = railway.AndThen(response, getURLByUUID)

		return response.record, response.err
	}
}
