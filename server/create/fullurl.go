package create

import (
	"errors"
)

var (
	FullURLDuplicatedErr = ObjectValueUniquenessError{pattern: "create: URL full value [%v] is not unique"}
)

type fullURL string

func (f fullURL) validateUniqueness(err error) error {
	if _, errR := repository.GetByFull(string(f)); errR == nil {
		return errors.Join(err, FullURLDuplicatedErr.From(f))
	}

	return err
}

func (f fullURL) Validate(err error) error {
	err = f.validateUniqueness(err)

	return err
}
