package redirect

import (
	"context"
	"embed"
	"github.com/NordGus/shrtnr/server/storage"
	"html/template"
)

var (
	environment string
	host        string

	//go:embed templates
	templates embed.FS

	helpers = template.FuncMap{
		"environment": func() string { return environment },
	}

	ctx        context.Context
	repository Repository
)

func Start(parentCtx context.Context, env string, redirectHost string) {
	ctx = parentCtx
	environment = env
	host = redirectHost
	repository = storage.GetURLRepository()
}
