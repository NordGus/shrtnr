package url

import (
	"github.com/NordGus/shrtnr/http/helpers"
	"html/template"
	"io/fs"
)

var (
	environment string
	views       *template.Template
)

func Start(env string, templates fs.FS) error {
	var err error

	environment = env

	views, err = template.New("applet").Funcs(helpers.Base).ParseFS(templates, "templates/url/applet.gohtml")
	if err != nil {
		return err
	}

	views, err = views.New("url").Funcs(helpers.Base).ParseFS(templates, "templates/url/url.gohtml")
	if err != nil {
		return err
	}

	views, err = views.New("search_form").Funcs(helpers.Base).ParseFS(templates, "templates/url/search_form.gohtml")
	if err != nil {
		return err
	}

	views, err = views.New("page").Funcs(helpers.Base).ParseFS(templates, "templates/url/page.gohtml")
	if err != nil {
		return err
	}

	views, err = views.New("search_results").Funcs(helpers.Base).ParseFS(templates, "templates/url/search_results.gohtml")
	if err != nil {
		return err
	}

	views, err = views.New("created").Funcs(helpers.Base).ParseFS(templates, "templates/url/created.gohtml")
	if err != nil {
		return err
	}

	views, err = views.New("form").Funcs(helpers.Base).ParseFS(templates, "templates/url/form.gohtml")
	if err != nil {
		return err
	}

	return err
}
