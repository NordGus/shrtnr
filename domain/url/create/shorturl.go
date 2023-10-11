package create

import (
	"errors"
)

const (
	shortURLLength = 8
)

var (
	ShortURLInvalidLengthErr = errors.New("create: shortURL is too long")
	ShortURLDuplicatedErr    = errors.New("create: shortURL is not unique")
)

type shortURL string

func (s shortURL) validateLength() error {
	if len(s) > shortURLLength {
		return ShortURLInvalidLengthErr
	}

	return nil
}

func (s shortURL) validateUniqueness() error {
	if _, err := repository.GetByShort(string(s)); err == nil {
		return ShortURLDuplicatedErr
	}

	return nil
}

func (s shortURL) Validate() error {
	return errors.Join(s.validateLength(), s.validateUniqueness())
}
