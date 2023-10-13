package redirect

import (
	"net/http"

	"github.com/NordGus/shrtnr/domain/shared/railway"
	"github.com/NordGus/shrtnr/domain/url"
)

func GetTarget(r *http.Request) (url.Target, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		response := extractShortFromPath(r)
		response = railway.AndThen(response, searchTarget)

		return response.target, response.err
	}
}
