package queue

import list "github.com/OddEer0/go-data-structure/list/double_list"

type IQueue[T any] interface {
	Unshift(item T)
	Shift() T
	Size() int
}

func NewQueue[T any]() IQueue[T] {
	return list.New[T]()
}
