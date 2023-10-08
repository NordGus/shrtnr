package redirect

import (
	"fmt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Host == host {
			// TODO: implement redirect logic

			err := fmt.Errorf("redirect: failed to redirect from %s: to be implemented", r.URL.RawPath)

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		next.ServeHTTP(w, r)
	})
}
