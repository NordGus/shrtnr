package create

import (
	"github.com/NordGus/shrtnr/domain/shared/response"
	"github.com/NordGus/shrtnr/domain/storage/url"
)

func AddURL(short string, full string) (url.URL, error) {
	select {
	case <-ctx.Done():
		return url.URL{}, ctx.Err()
	default:
		lock.Lock()
		defer lock.Unlock()

		resp := response.AndThen(buildUrl(short, full), validateUrl)
		resp = response.AndThen(resp, canBeAdded)
		resp = response.OnFailure(resp, deleteOldestUrl)
		resp = response.AndThen(resp, persistNewURl)
		resp = response.AndThen(resp, addUrlToQueue)

		return resp.record, resp.err
	}
}
