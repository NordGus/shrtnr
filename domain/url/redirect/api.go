package redirect

import (
	"net/http"

	"github.com/NordGus/shrtnr/domain/shared/railway"
)

func GetTarget(r *http.Request) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		resp := extractShortFromPath(r)
		resp = railway.AndThen(resp, searchFullURL)

		return resp.full, resp.err
	}
}
