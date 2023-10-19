package url

import "net/http"

type NewURLViewModel struct {
	ID     string
	UUID   string
	Target string
}

func AppletHandler(w http.ResponseWriter, _ *http.Request) {
	err := views.ExecuteTemplate(w, "applet", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewURLHandler(w http.ResponseWriter, _ *http.Request) {
	var vm NewURLViewModel

	err := views.ExecuteTemplate(w, "form", vm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateURLHandler(w http.ResponseWriter, r *http.Request) {}
