package create

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

func AddURL(entity entities.URL) (entities.URL, entities.URL, error) {
	select {
	case <-ctx.Done():
		return entities.URL{}, entities.URL{}, ctx.Err()
	default:
		response := newAddURLResponse(entity)
		response = railway.Then(response, validateUUIDUniqueness)
		response = railway.Then(response, validateTargetUniqueness)
		response = railway.AndThen(response, canBeAdded)
		response = railway.IfFailed(response, deleteOldestUrl)
		response = railway.AndThen(response, persistNewURl)
		response = railway.AndThen(response, addUrlToQueue)

		return response.record, response.oldRecord, response.err
	}
}
