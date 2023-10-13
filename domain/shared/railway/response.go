// Package railway is inspired in the concept of Railway Oriented Programming
// (https://blog.logrocket.com/what-is-railway-oriented-programming/) concept
// that I learned in my short stint at Job&Talent
package railway

type Response interface {
	Success() bool
}

// AndThen executes next if the previous Response was successful
func AndThen[T Response](r T, next func(in T) T) T {
	if r.Success() {
		return next(r)
	}

	return r
}

// OrThen executes next regardless if the previous Response was successful or not
func OrThen[T Response](r T, next func(in T) T) T {
	return next(r)
}

// OnFailure executes next if the previous Response wasn't successful
func OnFailure[T Response](r T, next func(in T) T) T {
	if r.Success() {
		return r
	}

	return next(r)
}
