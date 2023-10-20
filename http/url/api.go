package url

import (
	"github.com/NordGus/shrtnr/domain/url"
	"log"
	"net/http"
	"strconv"
	"time"
)

type NewURLViewModel struct {
	ID     string
	UUID   string
	Target string
}

type urlRecord struct {
	ID        string
	UUID      string
	Target    string
	CreatedAt time.Time
	DeletedAt time.Time
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

func (r urlRecord) SetDeletedAt(deletedAt time.Time) urlRecord {
	r.DeletedAt = deletedAt
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	rcrds, total, err := url.PaginateURLs(uint(page), rcrds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		URLs     []urlRecord
		PerPage  uint
		NextPage uint
	}{
		make([]urlRecord, total),
		10,
		uint(page) + 1,
	}

	copy(data.URLs, rcrds)

	err = views.ExecuteTemplate(w, "urls", data)
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

func CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	var (
		id     = r.FormValue("id")
		uuid   = r.FormValue("uuid")
		target = r.FormValue("target")

		vm   NewURLViewModel
		rcrd urlRecord
		err  error
	)

	rcrd, err = url.CreateURL(id, uuid, target, rcrd)
	if err != nil {
		log.Println(err)
		vm.ID = id
		vm.UUID = uuid
		vm.Target = target

		err = views.ExecuteTemplate(w, "form", vm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = views.ExecuteTemplate(w, "url", rcrd)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = views.ExecuteTemplate(w, "form", vm)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
