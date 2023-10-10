package helpers

import "html/template"

var (
	Base = template.FuncMap{
		"environment": func() string { return environment },
	}
)
