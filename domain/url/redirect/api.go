package redirect

import (
	"net/http"

	"github.com/NordGus/shrtnr/domain/shared/response"
)

func GetTarget(r *http.Request) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		resp := extractShortFromPath(r)
		resp = response.AndThen(resp, searchFullURL)

		return resp.full, resp.err
	}
}
