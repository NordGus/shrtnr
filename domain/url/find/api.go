package find

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

func PaginateURLs(page uint, perPage uint) ([]url.URL, error) {
	select {
	case <-ctx.Done():
		return []url.URL{}, ctx.Err()
	default:
		resp := buildPaginationSignal(page, perPage)
		resp = railway.AndThen(resp, getURLs)

		return resp.records, resp.err
	}
}

func GetURL(id uint) (url.URL, error) {
	select {
	case <-ctx.Done():
		return url.URL{}, ctx.Err()
	default:
		resp := buildGetSignal(id)
		resp = railway.AndThen(resp, getURL)

		return resp.record, resp.err
	}
}
