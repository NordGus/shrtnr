package redirect

import (
	"fmt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check possible problems with this implementation
		if r.Host != host {
			next.ServeHTTP(w, r)
			return
		}

		target, err := GetTarget(r)
		if err == nil {
			http.RedirectHandler(target.String(), http.StatusMovedPermanently).ServeHTTP(w, r)
			return
		}

		err = view.Execute(w, fmt.Sprintf("http://%s%s", r.Host, r.URL.Path))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
