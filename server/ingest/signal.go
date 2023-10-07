package ingest

type signal struct {
	new Url
	old Url
	err error
}

func (s signal) Error() error {
	return s.err
}
