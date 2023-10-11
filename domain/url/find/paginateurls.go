package find

import (
	"github.com/NordGus/shrtnr/domain/url/storage/url"
)

type paginationSignal struct {
	records []url.URL
	page    uint
	perPage uint
	err     error
}

func (s paginationSignal) Error() error {
	return s.err
}

func buildPaginationSignal(page uint, perPage uint) paginationSignal {
	return paginationSignal{page: page, perPage: perPage}
}

func getURLs(sig paginationSignal) paginationSignal {
	sig.records, sig.err = repository.GetAllInPage(sig.page, sig.perPage)

	return sig
}
