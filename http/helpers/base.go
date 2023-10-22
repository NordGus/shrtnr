package helpers

import (
	"fmt"
	"html/template"
	"time"
)

var (
	Base = template.FuncMap{
		"environment":      func() string { return environment },
		"formatDate":       func(t time.Time) string { return t.Format("01/02/2006 15:04:05") },
		"toRedirectionURL": func(uuid string) string { return fmt.Sprintf("%s/%s", redirectURL, uuid) },
	}
)
