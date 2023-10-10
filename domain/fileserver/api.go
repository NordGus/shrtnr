package fileserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes(r chi.Router) {
	r.Route("/content", func(cr chi.Router) {
		cr.Get("/dist", http.RedirectHandler("/content/dist/", http.StatusMovedPermanently).ServeHTTP)
		cr.Get("/dist/*", http.StripPrefix("/content", http.FileServer(http.FS(files))).ServeHTTP)

		cr.Group(func(pcr chi.Router) {
			pcr.Use(placeholderMiddleware)

			pcr.Get("/private", http.RedirectHandler("/content/private/", http.StatusMovedPermanently).ServeHTTP)
			pcr.Get("/private/*", http.StripPrefix("/content", http.FileServer(http.FS(files))).ServeHTTP)
		})
	})
}

func placeholderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("fileserver: please replace with authentication middleware when available")

		next.ServeHTTP(w, r)
	})
}
