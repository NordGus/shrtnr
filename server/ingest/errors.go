package ingest

import "fmt"

type ObjectValueValueError struct {
	got     any
	wanted  any
	pattern string
}

func (e *ObjectValueValueError) Error() string {
	return fmt.Sprintf(e.pattern, e.got, e.wanted)
}

func (e *ObjectValueValueError) From(got any, wanted any) *ObjectValueValueError {
	err := *e
	err.got = got
	err.wanted = wanted

	return &err
}

type ObjectValueUniquenessError struct {
	value   any
	pattern string
}

func (e *ObjectValueUniquenessError) Error() string {
	return fmt.Sprintf(e.pattern, e.value)
}

func (e *ObjectValueUniquenessError) From(value any) *ObjectValueUniquenessError {
	err := *e
	err.value = value

	return &err
}
