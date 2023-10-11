package find

import (
	"github.com/NordGus/shrtnr/domain/shared/response"
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

func PaginateURLs(page uint, perPage uint) ([]url.URL, error) {
	select {
	case <-ctx.Done():
		return []url.URL{}, ctx.Err()
	default:
		resp := buildPaginationSignal(page, perPage)
		resp = response.AndThen(resp, getURLs)

		return resp.records, resp.err
	}
}

func GetURL(id uint) (url.URL, error) {
	select {
	case <-ctx.Done():
		return url.URL{}, ctx.Err()
	default:
		resp := buildGetSignal(id)
		resp = response.AndThen(resp, getURL)

		return resp.record, resp.err
	}
}
