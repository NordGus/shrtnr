package create

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url"
)

func AddURL(entity url.URL) (url.URL, error) {
	select {
	case <-ctx.Done():
		return url.URL{}, ctx.Err()
	default:
		lock.Lock()
		defer lock.Unlock()

		resp := newAddURLResponse(entity)
		resp = railway.AndThen(resp, canBeAdded)
		resp = railway.OnFailure(resp, deleteOldestUrl)
		resp = railway.AndThen(resp, persistNewURl)
		resp = railway.AndThen(resp, addUrlToQueue)

		return resp.record, resp.err
	}
}
