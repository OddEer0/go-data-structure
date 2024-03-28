package listqueue

import (
	list "github.com/OddEer0/go-data-structure/list/double_list"
	queueInterface "github.com/OddEer0/go-data-structure/queue"
)

type (
	Queue[T any] interface {
		queueInterface.Queue[T]
		Copy() Queue[T]
	}

	queue[T any] struct {
		list list.List[T]
	}
)

func (q queue[T]) Size() int {
	return q.Size()
}

func (q queue[T]) Clear() {
	q.list.Clear()
}

func (q queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q queue[T]) String() string {
	return q.list.String()
}

func (q queue[T]) Unshift(item T) {
	q.list.Unshift(item)
}

func (q queue[T]) Shift() T {
	return q.list.Shift()
}

func (q queue[T]) Peek() T {
	return q.list.Head().Value()
}

func (q queue[T]) Copy() Queue[T] {
	return &queue[T]{list: q.list.Copy()}
}

// New TODO - add tests
func New[T any]() Queue[T] {
	return &queue[T]{list: list.New[T]()}
}
