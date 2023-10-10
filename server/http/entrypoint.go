package http

import (
	"embed"
	"errors"
	"github.com/NordGus/shrtnr/server/http/helpers"
	"html/template"
)

var (
	InitializationErr = errors.New("http: failed to initialize")

	environment string

	//go:embed templates
	templates embed.FS

	views *template.Template
)

func Start(env string) error {
	var err error

	// sub-packages initialization
	helpers.Start(env)

	environment = env

	// view initialization
	views, err = template.New("layout").Funcs(helpers.Base).ParseFS(templates, "templates/layout.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	views, err = views.New("app").Funcs(helpers.Base).ParseFS(templates, "templates/app.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	return err
}
