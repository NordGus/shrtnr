package create

import (
	"errors"
)

const (
	shortURLLength = 8
)

var (
	ShortURLInvalidLengthErr = ObjectValueValueError{pattern: "create: URL short value length invalid got [%v] wanted [%v]"}
	ShortURLDuplicatedErr    = ObjectValueUniquenessError{pattern: "create: URL short value [%v] is not unique"}
)

type shortURL string

func (s shortURL) validateLength(err error) error {
	if len(s) != shortURLLength {
		return errors.Join(err, ShortURLInvalidLengthErr.From(len(s), shortURLLength))
	}

	return err
}

func (s shortURL) validateUniqueness(err error) error {
	if _, errR := repository.GetByShort(string(s)); errR == nil {
		return errors.Join(err, ShortURLDuplicatedErr.From(s))
	}

	return err
}

func (s shortURL) Validate(err error) error {
	err = s.validateLength(err)
	err = s.validateUniqueness(err)

	return err
}
