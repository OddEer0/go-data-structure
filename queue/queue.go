package queue

import "github.com/OddEer0/go-data-structure/list"

type IQueue[T any] interface {
	PeekFirst() T
	Unshift(item T)
	Shift() T
	Size() int
}

func NewQueue[T any]() IQueue[T] {
	return list.NewLinkedList[T]()
}
