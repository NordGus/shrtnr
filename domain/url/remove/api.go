package remove

import (
	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

// DeleteURL can panic if deletion propagation fails. Because it means the system is corrupted and can't be trusted
func DeleteURL(id entities.ID) (entities.URL, error) {
	select {
	case <-ctx.Done():
		return entities.URL{}, ctx.Err()
	default:
		lock.Lock()
		defer lock.Unlock()

		response := newRemoveURLResponse(id)
		response = railway.AndThen(response, findRecord)
		response = railway.AndThen(response, deleteRecord)
		response = railway.AndThen(response, propagateDeletion)

		return response.record, response.err
	}
}
