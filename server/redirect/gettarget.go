package redirect

import (
	"errors"
	"net/http"
	"strings"
)

var (
	InvalidPathErr = errors.New("redirect: invalid path format")
)

type signal struct {
	short string
	full  string
	err   error
}

func (s signal) Error() error {
	return s.err
}

func extractShortFromPath(r *http.Request) signal {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 2 {
		return signal{err: InvalidPathErr}
	}

	return signal{short: path[1]}
}

func searchFullURL(sig signal) signal {
	record, err := repository.GetByShort(sig.short)
	if err != nil {
		sig.err = errors.Join(sig.err, err)
		return sig
	}

	sig.full = record.FullURL()

	return sig
}
