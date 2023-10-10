package http

import (
	"embed"
	"errors"
	"github.com/NordGus/shrtnr/server/http/helpers"
	"html/template"

	mnfst "github.com/NordGus/shrtnr/server/shared/manifest"
)

var (
	InitializationErr = errors.New("http: failed to initialize")

	environment string

	//go:embed templates
	templates embed.FS

	views    *template.Template
	manifest mnfst.Manifest
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

	manifest = mnfst.NewManifest(env)

	return err
}
