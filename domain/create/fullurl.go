package create

import (
	"errors"
)

var (
	FullURLDuplicatedErr = errors.New("create: fullURL is not unique")
)

type fullURL string

func (f fullURL) validateUniqueness() error {
	if _, err := repository.GetByFull(string(f)); err == nil {
		return FullURLDuplicatedErr
	}

	return nil
}

func (f fullURL) Validate() error {
	return errors.Join(f.validateUniqueness())
}
