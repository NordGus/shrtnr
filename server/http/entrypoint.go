package http

import "embed"

var (
	environment string

	//go:embed templates
	templates embed.FS
)

func Start(env string) {
	environment = env
}
