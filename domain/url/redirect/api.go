package redirect

import (
	"net/http"

	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url"
)

func GetTarget(r *http.Request) (url.URL, error) {
	select {
	case <-ctx.Done():
		return url.URL{}, ctx.Err()
	default:
		response := extractShortFromPath(r)
		response = railway.AndThen(response, getTargetRecord)

		return response.record, response.err
	}
}
