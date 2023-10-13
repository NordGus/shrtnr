package find

import (
	"github.com/NordGus/shrtnr/domain/url"
)

type paginateURLsResponse struct {
	records []url.URL
	page    uint
	perPage uint
	err     error
}

func (s paginateURLsResponse) Success() bool {
	return s.err == nil
}

func buildPaginationSignal(page uint, perPage uint) paginateURLsResponse {
	return paginateURLsResponse{page: page, perPage: perPage}
}

func getURLs(sig paginateURLsResponse) paginateURLsResponse {
	sig.records, sig.err = repository.GetAllInPage(sig.page, sig.perPage)

	return sig
}
