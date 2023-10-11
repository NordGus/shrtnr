// Package response is inspired in the concept of Railway Oriented Programming
// (https://blog.logrocket.com/what-is-railway-oriented-programming/) concept
// that I learned in my short stint at Job&Talent
package response

type Response interface {
	Error() error
}

func AndThen[T Response](r T, next func(in T) T) T {
	if r.Error() != nil {
		return r
	}

	return next(r)
}

func OnFailure[T Response](r T, next func(in T) T) T {
	if r.Error() == nil {
		return r
	}

	return next(r)
}