package http

import (
	"embed"
	"errors"
	"github.com/NordGus/shrtnr/http/url"
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

func Start(env string, redirectHost string) error {
	var err error

	// sub-packages initialization
	hlprs.Start(env, redirectHost)

	environment = env

	// view initialization
	views, err = template.New("layout").Funcs(hlprs.Base).ParseFS(templates, "templates/layout.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	manifest = mnfst.NewManifest(env)

	err = url.Start(environment, templates)
	if err != nil {
		return err
	}

	return err
}
