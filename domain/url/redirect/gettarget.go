package redirect

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/url/entities"
	"net/http"
	"strings"
)

var (
	InvalidPathErr = errors.New("redirect: invalid path format")
)

type getTargetResponse struct {
	uuid   entities.UUID
	record entities.URL
	err    error
}

func (s getTargetResponse) Success() bool {
	return s.err == nil
}

func extractShortFromPath(r *http.Request) getTargetResponse {
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 2 {
		return getTargetResponse{err: InvalidPathErr}
	}

	uuid, err := entities.NewUUID(path[1])
	if err != nil {
		return getTargetResponse{err: errors.Join(InvalidPathErr, err)}
	}

	return getTargetResponse{uuid: uuid}
}

func getTargetRecord(response getTargetResponse) getTargetResponse {
	record, err := repository.GetByUUID(response.uuid)
	if err != nil {
		response.err = errors.Join(response.err, err)
		return response
	}

	response.record = record

	return response
}
