package ingest

type URL struct {
	short shortURL
	full  fullURL
}

func (r *URL) Validate() error {
	err := r.short.Validate(nil)
	err = r.full.Validate(err)

	return err
}
