package queue

import "errors"

var (
	ErrQueueIsFull  = errors.New("queue: is full")
	ErrQueueIsEmpty = errors.New("queue: is empty")
)

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
		return ErrQueueIsFull
	}

	q.items = append(q.items, item)
	q.count++

	return nil
}

func (q *Queue[T]) Pop() (T, error) {
	var out T

	if q.count == 0 {
		return out, ErrQueueIsEmpty
	}

	out = q.items[0]
	q.items = q.items[1:]
	q.count--

	return out, nil
}
