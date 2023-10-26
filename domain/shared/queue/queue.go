package queue

import "errors"

var (
	IsFullErr       = errors.New("queue: is full")
	IsEmptyErr      = errors.New("queue: is empty")
	ItemNotFoundErr = errors.New("queue: item not found")
)

type CompareFunc[T any] func(i T, j T) bool

type Queue[T any] struct {
	items []T
	count uint
	size  uint
}

func NewQueue[T any](size uint) Queue[T] {
	return Queue[T]{
		items: make([]T, 0, size),
		count: 0,
		size:  size,
	}
}

func (q *Queue[T]) Push(item T) error {
	if q.count == q.size {
		return IsFullErr
	}

	q.items = append(q.items, item)
	q.count++

	return nil
}

// FindAndRemoveBy searches an item in the Queue and removes it. If it doesn't find it returns ItemNotFoundErr.
func (q *Queue[T]) FindAndRemoveBy(item T, is CompareFunc[T]) error {
	for i := 0; i < len(q.items); i++ {
		if !is(item, q.items[i]) {
			continue
		}

		items := make([]T, 0, q.size)
		items = append(items, q.items[:i]...)
		items = append(items, q.items[i+1:]...)

		q.items = items
		q.count--

		return nil
	}

	return ItemNotFoundErr
}

func (q *Queue[T]) Pop() (T, error) {
	var out T

	if q.count == 0 {
		return out, IsEmptyErr
	}

	out = q.items[0]
	q.items = q.items[1:]
	q.count--

	return out, nil
}

func (q *Queue[T]) Peek() (T, error) {
	var peek T

	if q.count == 0 {
		return peek, IsEmptyErr
	}

	return q.items[0], nil
}

func (q *Queue[T]) IsFull() bool {
	return q.count == q.size
}

func (q *Queue[T]) Size() uint {
	return q.count
}
