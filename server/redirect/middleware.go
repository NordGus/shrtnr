package redirect

import (
	"errors"
	"html/template"
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
			http.RedirectHandler(target, http.StatusMovedPermanently).ServeHTTP(w, r)
			return
		}

		renderErr := renderErrorView(w, err)
		if renderErr != nil {
			http.Error(w, renderErr.Error(), http.StatusInternalServerError)
		}
	})
}

func renderErrorView(w http.ResponseWriter, redirectErr error) error {
	tmpl, err := template.New("error").Funcs(helpers).ParseFS(templates, "templates/error.gohtml")
	if err != nil {
		return errors.Join(redirectErr, err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		return errors.Join(redirectErr, err)
	}

	return nil
}
