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

		response := newAddURLResponse(entity)
		response = railway.OrThen(response, validateUUIDUniqueness)
		response = railway.OrThen(response, validateTargetUniqueness)
		response = railway.AndThen(response, canBeAdded)
		response = railway.OnFailure(response, deleteOldestUrl)
		response = railway.AndThen(response, persistNewURl)
		response = railway.AndThen(response, addUrlToQueue)

		return response.record, response.err
	}
}
