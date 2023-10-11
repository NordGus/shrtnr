package redirect

import (
	"context"
	"embed"
	"errors"
	"html/template"

	hlprs "github.com/NordGus/shrtnr/domain/url/redirect/helpers"
	"github.com/NordGus/shrtnr/domain/url/storage"
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

	hlprs.Start(environment)

	view, err = template.New("error").Funcs(hlprs.Base).ParseFS(templates, "templates/error.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	return err
}
