package http

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func Routes(r chi.Router) {
	log.Println("setting up http routes for environment:", environment)

	r.Get("/", baseHandler)
	r.Get("/app", appHandler)
}

func baseHandler(w http.ResponseWriter, _ *http.Request) {
	err := views.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func appHandler(w http.ResponseWriter, _ *http.Request) {
	err := views.ExecuteTemplate(w, "app", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
