package redirect

import (
	"fmt"
	"net/http"

	"github.com/NordGus/shrtnr/domain/url"

	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/r/{uuid}", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	var (
		uuid   = chi.URLParam(r, "uuid")
		target Target
	)

	target, err := url.FindURLByUUID(uuid, target)
	if err == nil {
		http.RedirectHandler(target.redirectTo, http.StatusMovedPermanently).ServeHTTP(w, r)
		return
	}

	err = view.Execute(w, fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
