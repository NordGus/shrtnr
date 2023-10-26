package url

import (
	"fmt"
	"github.com/NordGus/shrtnr/domain/url"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

type NewURLForm struct {
	ID     string
	UUID   string
	Target string
}

type SearchURLForm struct {
	Term string
}

type urlRecord struct {
	ID        string
	UUID      string
	Target    string
	CreatedAt time.Time
}

func (r urlRecord) SetID(id string) urlRecord {
	r.ID = id
	return r
}

func (r urlRecord) SetUUID(uuid string) urlRecord {
	r.UUID = uuid
	return r
}

func (r urlRecord) SetTarget(target string) urlRecord {
	r.Target = target
	return r
}

func (r urlRecord) SetCreatedAt(createdAt time.Time) urlRecord {
	r.CreatedAt = createdAt
	return r
}

func AppletHandler(w http.ResponseWriter, _ *http.Request) {
	err := views.ExecuteTemplate(w, "applet", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetURLsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		rcrds        = make([]urlRecord, 10)
		page  uint64 = 1
		err   error
	)

	if r.URL.Query().Has("page") {
		page, err = strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
		if err != nil {
			err = views.ExecuteTemplate(w, "error_toast", err.Error())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	rcrds, total, err := url.PaginateURLs(uint(page), rcrds)
	if err != nil {
		err = views.ExecuteTemplate(w, "error_toast", err.Error())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	data := struct {
		URLs       []urlRecord
		PerPage    uint
		NextPage   uint
		SearchForm SearchURLForm
	}{
		make([]urlRecord, total),
		10,
		uint(page) + 1,
		SearchURLForm{},
	}

	copy(data.URLs, rcrds)

	err = views.ExecuteTemplate(w, "page", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewURLHandler(w http.ResponseWriter, _ *http.Request) {
	var vm NewURLForm

	err := views.ExecuteTemplate(w, "form", vm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	var (
		id     = r.FormValue("id")
		uuid   = r.FormValue("uuid")
		target = r.FormValue("target")

		vm      NewURLForm
		rcrd    urlRecord
		oldrcrd urlRecord
		err     error
	)

	rcrd, oldrcrd, err = url.CreateURL(id, uuid, target, rcrd, oldrcrd)
	if err != nil {
		log.Println(err)

		vm.ID = id
		vm.UUID = uuid
		vm.Target = target

		err = views.ExecuteTemplate(w, "error_toast", err.Error())
		if err != nil {
			log.Println(err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if err == nil {
		err = views.ExecuteTemplate(w, "created", rcrd)
		if err != nil {
			log.Println(err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = views.ExecuteTemplate(w, "success_toast", fmt.Sprintf("%s added to system", rcrd.Target))
		if err != nil {
			log.Println(err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err == nil && oldrcrd.UUID != "" {
		err = views.ExecuteTemplate(w, "deleted", oldrcrd)
		if err != nil {
			log.Println(err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = views.ExecuteTemplate(w, "success_toast", fmt.Sprintf("%s removed from system because of overflow", oldrcrd.Target))
		if err != nil {
			log.Println(err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = views.ExecuteTemplate(w, "form", vm)
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteURLHandler(w http.ResponseWriter, r *http.Request) {
	var (
		id = chi.URLParam(r, "id")

		rcrd urlRecord
	)

	rcrd, err := url.RemoveURL(id, rcrd)
	if err != nil {
		err = views.ExecuteTemplate(w, "error_toast", err.Error())
		if err != nil {
			log.Println(err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	err = views.ExecuteTemplate(w, "deleted", rcrd)
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = views.ExecuteTemplate(w, "success_toast", fmt.Sprintf("%s removed from system", rcrd.Target))
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		rcrds = make([]urlRecord, 0)
		term  = r.FormValue("term")
		err   error
	)

	if term == "" {
		w.WriteHeader(http.StatusOK)
		return
	}

	rcrds, err = url.SearchURLsBy(term, rcrds)
	if err != nil {
		log.Println(err)

		err = views.ExecuteTemplate(w, "error_toast", err.Error())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	data := struct {
		URLs []urlRecord
	}{
		rcrds,
	}

	err = views.ExecuteTemplate(w, "search_results", data)
	if err != nil {
		log.Println(err)

		err = views.ExecuteTemplate(w, "error_toast", err.Error())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
