package create

import "errors"

type URL struct {
	short shortURL
	full  fullURL
}

func (r *URL) Validate() error {
	return errors.Join(r.short.Validate(), r.full.Validate())
}
