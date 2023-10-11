package http

import (
	"embed"
	"errors"
	"html/template"

	hlprs "github.com/NordGus/shrtnr/http/helpers"
	mnfst "github.com/NordGus/shrtnr/http/manifest"
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
	hlprs.Start(env)

	environment = env

	// view initialization
	views, err = template.New("layout").Funcs(hlprs.Base).ParseFS(templates, "templates/layout.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	views, err = views.New("app").Funcs(hlprs.Base).ParseFS(templates, "templates/app.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	manifest = mnfst.NewManifest(env)

	return err
}
