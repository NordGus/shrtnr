package redirect

import (
	"embed"
	"errors"
	"github.com/NordGus/shrtnr/redirect/helpers"
	"html/template"
)

var (
	InitializationErr = errors.New("redirect: failed to initialize")

	//go:embed templates
	templates embed.FS
	view      *template.Template
)

func Start(env string) error {
	var err error

	helpers.Start(env)

	view, err = template.New("error").Funcs(helpers.Base).ParseFS(templates, "templates/error.gohtml")
	if err != nil {
		return errors.Join(InitializationErr, err)
	}

	return err
}
