package redirect

import (
	"context"
	"embed"
	"errors"
	"github.com/NordGus/shrtnr/server/redirect/helpers"
	"github.com/NordGus/shrtnr/server/storage"
	"html/template"
)

var (
	InitializationErr = errors.New("redirect: failed to initialize")

	environment string
	host        string

	//go:embed templates
	templates embed.FS
	view      *template.Template

	ctx        context.Context
	repository Repository
)

func Start(parentCtx context.Context, env string, redirectHost string) error {
	var err error

	ctx = parentCtx
	environment = env
	host = redirectHost
	repository = storage.GetURLRepository()

	helpers.Start(environment)

	view, err = template.New("error").Funcs(helpers.Base).ParseFS(templates, "templates/error.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	return err
}
