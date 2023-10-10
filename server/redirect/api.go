package redirect

import (
	"github.com/NordGus/shrtnr/server/shared/response"
	"net/http"
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
