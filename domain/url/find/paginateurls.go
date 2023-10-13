package find

import (
	"github.com/NordGus/shrtnr/domain/url/entities"
)

type paginateURLsResponse struct {
	records []entities.URL
	page    uint
	perPage uint
	err     error
}

func (s paginateURLsResponse) Success() bool {
	return s.err == nil
}

func buildPaginateURLsResponse(page uint, perPage uint) paginateURLsResponse {
	return paginateURLsResponse{page: page, perPage: perPage}
}

func getURLs(response paginateURLsResponse) paginateURLsResponse {
	response.records, response.err = repository.GetAllInPage(response.page, response.perPage)

	return response
}
