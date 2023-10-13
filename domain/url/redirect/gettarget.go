package redirect

import (
	"errors"
	"github.com/NordGus/shrtnr/domain/url"
	"net/http"
	"strings"
)

var (
	InvalidPathErr = errors.New("redirect: invalid path format")
)

type getTargetResponse struct {
	uuid   url.UUID
	target url.Target
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

	uuid, err := url.NewUUID(path[1])
	if err != nil {
		return getTargetResponse{err: errors.Join(InvalidPathErr, err)}
	}

	return getTargetResponse{uuid: uuid}
}

func searchTarget(response getTargetResponse) getTargetResponse {
	record, err := repository.GetByUUID(response.uuid)
	if err != nil {
		response.err = errors.Join(response.err, err)
		return response
	}

	response.target = record.Target

	return response
}
