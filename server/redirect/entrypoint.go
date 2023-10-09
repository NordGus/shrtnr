package redirect

import (
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

	repository Repository
)

func Start(env string, redirectHost string) {
	environment = env
	host = redirectHost
	repository = storage.GetURLRepository()
}
