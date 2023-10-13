package find

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url"
)

func PaginateURLs(page uint, perPage uint) ([]url.URL, error) {
	select {
	case <-ctx.Done():
		return []url.URL{}, ctx.Err()
	default:
		response := buildPaginateURLsResponse(page, perPage)
		response = railway.AndThen(response, getURLs)

		return response.records, response.err
	}
}

func GetURL(id url.ID) (url.URL, error) {
	select {
	case <-ctx.Done():
		return url.URL{}, ctx.Err()
	default:
		response := buildGetURLResponse(id)
		response = railway.AndThen(response, getURL)

		return response.record, response.err
	}
}
