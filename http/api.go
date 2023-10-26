package http

import (
	"encoding/json"
	"github.com/NordGus/shrtnr/http/url"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Routes(r chi.Router) {
	log.Println("setting up http routes for environment:", environment)

	r.Get("/manifest.json", manifestHandler)
	r.Get("/", baseHandler)
	r.Get("/url", url.AppletHandler)
	r.Get("/urls", url.GetURLsHandler)
	r.Get("/url/new", url.NewURLHandler)
	r.Put("/url/create", url.CreateURLHandler)
	r.Post("/url/search", url.GetSearchResultsHandler)
	r.Delete("/url/{id}", url.DeleteURLHandler)
}

func manifestHandler(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(manifest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func baseHandler(w http.ResponseWriter, _ *http.Request) {
	err := views.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
