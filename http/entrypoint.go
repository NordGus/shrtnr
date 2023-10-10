package http

import (
	"embed"
	"errors"
	helpers2 "github.com/NordGus/shrtnr/http/helpers"
	"html/template"

	mnfst "github.com/NordGus/shrtnr/domain/shared/manifest"
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
	helpers2.Start(env)

	environment = env

	// view initialization
	views, err = template.New("layout").Funcs(helpers2.Base).ParseFS(templates, "templates/layout.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	views, err = views.New("app").Funcs(helpers2.Base).ParseFS(templates, "templates/app.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	manifest = mnfst.NewManifest(env)

	return err
}
