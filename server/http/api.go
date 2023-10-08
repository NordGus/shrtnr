package http

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func Routes(r chi.Router) {
	log.Println("setting up http routes for environment:", environment)

	r.Get("/", baseHandler)
	r.Get("/app", appHandler)
}

func baseHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.New("layout").ParseFS(templates, "templates/layout.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func appHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.New("app").ParseFS(templates, "templates/app.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
