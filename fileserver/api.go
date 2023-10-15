package fileserver

import (
	"embed"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

var (
	// files most only contain 2 directories one called dist (for public static files)
	// and another called private (for private static files)
	//go:embed dist private
	files embed.FS
)

func PublicRoutes(r chi.Router) {
	r.Get("/dist", http.RedirectHandler("/dist/", http.StatusMovedPermanently).ServeHTTP)
	r.Get("/dist/*", http.FileServer(http.FS(files)).ServeHTTP)
}

func PrivateRoutes(r chi.Router) {
	r.Group(func(pr chi.Router) {
		pr.Use(placeholderMiddleware)

		pr.Get("/private", http.RedirectHandler("/private/", http.StatusMovedPermanently).ServeHTTP)
		pr.Get("/private/*", http.FileServer(http.FS(files)).ServeHTTP)
	})
}

func placeholderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("fileserver: please replace with authentication middleware when available")

		next.ServeHTTP(w, r)
	})
}
