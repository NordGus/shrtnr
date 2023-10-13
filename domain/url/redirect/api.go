package redirect

import (
	"net/http"

	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url/entities"
)

func GetTarget(r *http.Request) (entities.URL, error) {
	select {
	case <-ctx.Done():
		return entities.URL{}, ctx.Err()
	default:
		response := extractShortFromPath(r)
		response = railway.AndThen(response, getTargetRecord)

		return response.record, response.err
	}
}
